package main

import (
	"math/rand"
	"time"
)

func chooseRandomlyInitialTrigram(trigramMap TrigramMap) Trigram {

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

	// I know this is an extremly bad practice, but I made it this way to be simpler for this function and its returns.
	// In production I wouldn't do this.
	if word1 == "" || word2 == "" || word3 == "" {
		panic("error in random generation, something went wrong")
	}

	return Trigram{word1, word2, word3}
}

func chooseNextWord(wordFreqs map[string]int) string {

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

	panic("No word was generated!")
}

func main() {
	server := NewServer(NewTrigramStore(chooseRandomlyInitialTrigram, chooseNextWord))
	server.Run()
}
