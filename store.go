package main

// TrigramStore represents the storage of trigrams found until now.
type TrigramStore struct {
	trigrams map[string]map[string]map[string]int
}
