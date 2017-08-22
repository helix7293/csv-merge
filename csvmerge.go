package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Result struct {
	creatorId        string
	creatorFirstName string
	creatorLastName  string
	pagePath         string
	lastTraffic      string
}

func (r Result) Columns() []string {
	return []string{r.creatorId, r.creatorFirstName, r.creatorLastName, r.pagePath, r.lastTraffic}
}

func main() {
	usercsv, _ := os.Open("QueryResults.csv")
	segmentcsv, _ := os.Open("totalLastSeen.csv")

	defer usercsv.Close()
	defer segmentcsv.Close()

	reader := csv.NewReader(bufio.NewReader(usercsv))

	csvfile := map[string]*Result{}

	for {
		record, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		csvfile[record[3]] = &Result{creatorId: record[0], creatorFirstName: record[1], creatorLastName: record[2], pagePath: record[3]}
	}

	reader = csv.NewReader(bufio.NewReader(segmentcsv))
	for {
		record, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		pagePath := record[0]
		lastTraffic := record[1]

		result, ok := csvfile[pagePath]

		if ok {
			result.lastTraffic = lastTraffic
		}

	}

	fmt.Println("Id,First Name,Last Name,Page Path,Last Traffic")
	for _, result := range csvfile {
		fmt.Println(strings.Join(result.Columns(), ","))
	}
}
