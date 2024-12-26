package main

import (
	"encoding/json"
	"errors"
	"strings"
)

type Formatter interface {
	Format([][]string) (string, error)
}

type CSVFormatter struct {
	ReturnHeaders bool
}

func (cf *CSVFormatter) Format(records [][]string) (string, error) {
	var startIndex = 1
	if cf.ReturnHeaders {
		startIndex = 0
	}

	var lines []string

	for _, record := range records[startIndex:] {
		lines = append(lines, strings.Join(record, ","))
	}

	return strings.Join(lines, "\n"), nil
}

type JSONOutput struct {
	Sentence string
	Word     string
}

type JSONFormatter struct{}

func (jf *JSONFormatter) Format(records [][]string) (string, error) {

	var lines []JSONOutput
	for _, record := range records[1:] {
		lines = append(lines, JSONOutput{
			Sentence: record[0],
			Word:     record[1],
		})
	}

	j, err := json.Marshal(lines)
	return string(j), err
}

func newFormatter(format string) (Formatter, error) {
	switch format {
	case "csv":
		return &CSVFormatter{ReturnHeaders: true}, nil
	case "json":
		return &JSONFormatter{}, nil
	default:
		return nil, errors.New("Invalid format")
	}
}
