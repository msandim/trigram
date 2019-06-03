package main

import (
	"errors"
	"strings"
)

func learnText(store *TrigramStore, text string) error {

	// Was thinking of a go routine that does all of this, but it's probably better not because we want to know if the learning process was valid:

	// Get trigrams:
	trigrams, err := parseTrigrams(text)

	if err != nil {
		return err
	}

	for _, trigram := range trigrams {
		store.AddTrigram(trigram)
	}

	return nil
}

func parseTrigrams(text string) ([]Trigram, error) {
	words := strings.Split(text, " ")

	if len(words) < 3 {
		return nil, errors.New("text to learn has less than 3 words")
	}

	var trigrams []Trigram

	for i := 0; i < len(words)-2; i++ {
		trigram := Trigram{words[i], words[i+1], words[i+2]}
		trigrams = append(trigrams, trigram)
	}

	return trigrams, nil
}
