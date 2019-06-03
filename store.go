package main

import (
	"sync"
)

// Trigram represents a sequence of 3 strings
type Trigram [3]string

// TrigramStore represents the storage of trigrams found until now.
type TrigramStore struct {
	trigrams map[string]map[string]map[string]int
	mutex    *sync.Mutex
}

type trigramEnd struct {
	elem string
	freq int
}

// NewTrigramStore creates a new TrigramStore.
func NewTrigramStore() *TrigramStore {
	var store TrigramStore
	store.trigrams = make(map[string]map[string]map[string]int)
	store.mutex = &sync.Mutex{}
	return &store
}

// AddTrigram adds a trigram to the store, increasing its "popularity" if it's already present in the store.
func (store *TrigramStore) AddTrigram(trigram Trigram) {

	store.mutex.Lock()

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

	store.mutex.Unlock()
}

// MakeText generates a random text with the trigrams present in the store.
func (store *TrigramStore) MakeText() string {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	// TODO: generate text

	return "lol"
}

func (store *TrigramStore) getTrigramFreq(trigram [3]string) int {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	return store.trigrams[trigram[0]][trigram[1]][trigram[2]]
}
