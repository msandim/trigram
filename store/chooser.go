package store

import (
	"fmt"
	"math/rand"
	"time"
)

// Chooser is structure that knows how to choose the next trigram/word to use in a given moment.
type Chooser interface {
	// ChooseInitialTrigram chooses the initial trigram to start a text with, given a TrigramMap of available trigrams.
	ChooseInitialTrigram(availableTrigrams TrigramMap) Trigram
	// ChooseNextWord chooses the next word to use within a text, given the frequencies of each possible word to be used at this point.
	ChooseNextWord(possibleWords map[string]int) string
}

// RandomChooser implements a Chooser that makes random decisions.
type RandomChooser struct{}

// ChooseInitialTrigram chooses randomly a trigram to start the text.
func (c *RandomChooser) ChooseInitialTrigram(trigramMap TrigramMap) Trigram {
	rand.Seed(time.Now().UnixNano())

	var word1 string
	r1 := rand.Intn(len(trigramMap))
	for word := range trigramMap {
		r1--
		if r1 <= 0 {
			word1 = word
			break
		}
	}

	var word2 string
	r2 := rand.Intn(len(trigramMap[word1]))
	for word := range trigramMap[word1] {
		r2--
		if r2 <= 0 {
			word2 = word
			break
		}
	}

	var word3 string
	r3 := rand.Intn(len(trigramMap[word1][word2]))
	for word := range trigramMap[word1][word2] {
		r3--
		if r3 <= 0 {
			word3 = word
			break
		}
	}

	if word1 == "" || word2 == "" || word3 == "" {
		fmt.Println("WARNING: Could not randomly choose initial trigram to make text. Will use a trigram made of empty strings.")
	}

	return Trigram{word1, word2, word3}
}

// ChooseNextWord chooses randomly the next word to complete the text with.
// This random selection takes into account the frequencies of the sequencenin the learned texts.
func (c *RandomChooser) ChooseNextWord(wordFreqs map[string]int) string {

	// Count total frequencies:
	totalFreqs := 0
	for _, v := range wordFreqs {
		totalFreqs += v
	}

	rand.Seed(time.Now().UnixNano())

	partialFreq := rand.Intn(totalFreqs)
	for word, freq := range wordFreqs {
		partialFreq -= freq
		if partialFreq <= 0 {
			return word
		}
	}

	fmt.Println("WARNING: Could not choose the next word. Will use an empty string as the next word.")
	return ""
}
