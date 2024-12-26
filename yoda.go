package main

import (
	"math/rand"
	"strings"
)

func ShuffleString(value, delimiter string) string {
	words := strings.Split(value, delimiter)

	for i := range words {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, delimiter)
}

type Yoda interface {
	Shuffle() string
}

type SentenceYoda struct {
	Sentence string
}

func (sy *SentenceYoda) Shuffle() string {
	return ShuffleString(sy.Sentence, " ")
}

type WordYoda struct {
	Sentence string
	Word     string
}

func (wy *WordYoda) Shuffle() string {
	words := strings.Split(wy.Sentence, " ")
	for i := range words {
		if word := words[i]; word == wy.Word {
			words[i] = ShuffleString(word, "")
		}
	}
	return strings.Join(words, " ")
}

type AllYoda struct {
	Sentence string
}

func (ay *AllYoda) Shuffle() string {
	words := strings.Split(ay.Sentence, " ")
	for i, word := range words {
		words[i] = ShuffleString(word, "")
	}
	return strings.Join(words, " ")
}

func newYoda(sentence, word string) Yoda {
	switch word {
	case "":
		return &SentenceYoda{Sentence: sentence}
	case "all":
		return &AllYoda{Sentence: sentence}
	default:
		return &WordYoda{Sentence: sentence, Word: word}
	}
}
