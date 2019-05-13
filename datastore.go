package main

// WordOccurance is
type WordOccurance struct {
	word  string
	count int
}

// DataStore is
type DataStore struct {
	trigrams map[[2]string]WordOccurance
}
