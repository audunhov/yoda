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
Example sentence to shuffle,
Example sentence to shuffle,shuffle
Example sentence to shuffle,all
```

Meaning:
- Not having a word means the program shuffles the words
- Having a word means the program shuffles the letters in that word
- Having the word 'all' means the program shuffles the letters in all words

> Note that the program assumes the first line of the csv file is headers and will be skipped.

The program accepts the following parameters:
- file: path to CSV input file
- format: output format. Currently only CSV and JSON are supported
