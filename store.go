package main

import (
	"strings"
	"sync"
)

// Trigram represents a sequence of 3 strings
type Trigram [3]string

// TrigramMap is...
type TrigramMap map[string]map[string]map[string]int

// TrigramStore represents the storage of trigrams found until now.
type TrigramStore struct {
	trigrams              TrigramMap
	mutex                 *sync.Mutex
	initialTrigramChooser func(TrigramMap) Trigram
	nextWordChooser       func(map[string]int) string
}

// NewTrigramStore creates a new TrigramStore.
func NewTrigramStore(initialTrigramChooser func(TrigramMap) Trigram, nextWordChooser func(map[string]int) string) *TrigramStore {
	var store TrigramStore
	store.trigrams = make(map[string]map[string]map[string]int)
	store.mutex = &sync.Mutex{}
	store.initialTrigramChooser = initialTrigramChooser
	store.nextWordChooser = nextWordChooser
	return &store
}

// AddTrigram adds a trigram to the store, increasing its "popularity" if it's already present in the store.
func (store *TrigramStore) AddTrigram(trigram Trigram) {

	store.mutex.Lock()
	defer store.mutex.Unlock()

	elem0 := store.trigrams

	elem1, ok0 := elem0[trigram[0]]

	if !ok0 {
		elem1 = make(map[string]map[string]int)
		elem0[trigram[0]] = elem1
	}

	elem2, ok1 := elem1[trigram[1]]

	if !ok1 {
		elem2 = make(map[string]int)
		elem1[trigram[1]] = elem2
	}

	elem2[trigram[2]]++
}

// MakeText generates a random text with the trigrams present in the store.
func (store *TrigramStore) MakeText() string {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	var text []string
	var last2Words [2]string

	// Make a text with 10 trigrams:
	for i := 0; i < 10; i++ {
		if len(text) > 0 {
			// Choose the next word:
			possibleNextWords := store.trigrams[last2Words[0]][last2Words[1]]
			nextWord := store.nextWordChooser(possibleNextWords)
			text = append(text, nextWord)

			// Update the last 2 words:
			last2Words[0] = last2Words[1]
			last2Words[1] = nextWord
		} else {
			// Choose a random trigram:
			trigram := store.initialTrigramChooser(store.trigrams)
			text = append(text, trigram[:]...)

			// Update the last 2 words:
			last2Words[0] = trigram[1]
			last2Words[1] = trigram[2]
		}
	}

	return strings.Join(text, " ")
}

func (store *TrigramStore) getTrigramFreq(trigram Trigram) int {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	return store.trigrams[trigram[0]][trigram[1]][trigram[2]]
}
