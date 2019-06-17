package store

import (
	"strings"
	"sync"
)

// Trigram represents a sequence of 3 strings
type Trigram [3]string

// TrigramMap is a 3-dimensional map which represents the frequency of each trigram.
type TrigramMap map[string]map[string]map[string]int

// TrigramStore is
type TrigramStore interface {
	AddTrigram(trigram Trigram)
	MakeText() string
}

// TrigramMapStore represents the storage of trigrams found until now.
// It basically encapsulates a TrigramMap with a mutex and some functions which can be set to perform searches on the TrigramMap.
type TrigramMapStore struct {
	trigrams TrigramMap  // Check documentation of TrigramMap above.
	mutex    *sync.Mutex // Mutex to control accesses to the TrigramMap
	chooser  Chooser
}

// NewMapTrigramStore creates a new TrigramStore.
func NewMapTrigramStore(chooser Chooser) *TrigramMapStore {
	var store TrigramMapStore
	store.trigrams = make(map[string]map[string]map[string]int)
	store.mutex = &sync.Mutex{}
	store.chooser = chooser
	return &store
}

// AddTrigram adds a trigram to the store, increasing its "frequency" if it's already present in the store.
func (store *TrigramMapStore) AddTrigram(trigram Trigram) {

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
func (store *TrigramMapStore) MakeText() string {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	var text []string
	var last2Words [2]string

	if len(store.trigrams) == 0 {
		return ""
	}

	// Make a text with 100 trigrams maximum:
	for i := 0; i < 100; i++ {
		if len(text) > 0 {
			// Choose the next word, except if we encountered a path with zero possibilities for the next word.
			possibleNextWords := store.trigrams[last2Words[0]][last2Words[1]]
			if len(possibleNextWords) == 0 {
				break
			}
			nextWord := store.chooser.ChooseNextWord(possibleNextWords)
			text = append(text, nextWord)

			// Update the last 2 words:
			last2Words[0] = last2Words[1]
			last2Words[1] = nextWord
		} else {
			// Choose a random trigram to start:
			trigram := store.chooser.ChooseInitialTrigram(store.trigrams)
			text = append(text, trigram[:]...)

			// Update the last 2 words:
			last2Words[0] = trigram[1]
			last2Words[1] = trigram[2]
		}
	}

	return strings.Join(text, " ")
}

func (store *TrigramMapStore) getTrigramFreq(trigram Trigram) int {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	return store.trigrams[trigram[0]][trigram[1]][trigram[2]]
}
