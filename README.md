# Yoda

Simple program to shuffle words in sentences, or letters in a single word, or letters in all words

## Installation

```
go get github.com/audunhov/yoda@latest
```

## Usage

```
cat sentences.csv | yoda
```

The format of the CSV file should be:
```csv
Sentence,Word
Will shuffle the words but not the letters,
Will shuffle the word 'shuffle' but not the words,shuffle
Will shuffle the letters of all words but not the words,all
```

> Note that the program assumes the first line of the csv file is headers and will be skipped.

The program accepts the following parameters:
- file: path to CSV input file
- format: output format. Currently only CSV and JSON are supported
