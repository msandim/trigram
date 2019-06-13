package store

import (
	"sort"
	"testing"
)

func TestAddTrigram(t *testing.T) {

	store := NewTrigramStore(&RandomChooser{})
	var trigram [3]string
	var n int

	trigram = [3]string{"a", "b", "c"}
	store.AddTrigram(trigram)
	n = store.getTrigramFreq(trigram)

	if n != 1 {
		t.Fatalf("Trigram frequency is wrong. Got %d", n)
	}

	trigram = [3]string{"a", "b", "d"}
	store.AddTrigram(trigram)
	n = store.getTrigramFreq(trigram)

	if n != 1 {
		t.Fatalf("Trigram frequency is wrong. Got %d", n)
	}

	trigram = [3]string{"a", "b", "d"}
	store.AddTrigram(trigram)
	n = store.getTrigramFreq(trigram)

	if n != 2 {
		t.Fatalf("Trigram frequency is wrong. Got %d", n)
	}

	trigram = [3]string{"c", "b", "d"}
	store.AddTrigram(trigram)
	n = store.getTrigramFreq(trigram)

	if n != 1 {
		t.Fatalf("Trigram frequency is wrong. Got %d", n)
	}

	trigram = [3]string{"c", "b", "d"}
	store.AddTrigram(trigram)
	n = store.getTrigramFreq(trigram)

	if n != 2 {
		t.Fatalf("Trigram frequency is wrong. Got %d", n)
	}
}

type TestChooser struct{}

func (c *TestChooser) ChooseInitialTrigram(availableTrigrams TrigramMap) Trigram {
	return Trigram{"word1", "word2", "word3"}
}

func (c *TestChooser) ChooseNextWord(possibleWords map[string]int) string {
	var words []string

	for w := range possibleWords {
		words = append(words, w)
	}

	sort.Strings(words)
	return words[0]
}

func TestMakeText(t *testing.T) {

}
