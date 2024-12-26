package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "", "input file")
	format := flag.String("format", "csv", "output format")
	flag.Parse()

	formatter, err := newFormatter(*format)

	if err != nil {
		log.Fatal("Could not get formatter:", err)
	}

	var r io.Reader

	if filename == nil || *filename == "" {
		r = os.Stdin
	} else {
		r, err = os.Open(*filename)
		if err != nil {
			log.Fatal("Could not read from file,", err)
		}
	}

	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for _, record := range records {
		sentence, word := record[0], record[1]
		yoda := newYoda(sentence, word)
		shuffled := yoda.Shuffle()
		record[0] = shuffled
	}

	formatted, err := formatter.Format(records)

	if err != nil {
		log.Fatal("Could not format output:", err)
	}

	fmt.Println(formatted)

}
