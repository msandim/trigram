package store

import (
	"sort"
	"testing"
)

func TestAddTrigram(t *testing.T) {

	store := NewMapTrigramStore(&TestChooser{})
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

func TestMakeText(t *testing.T) {

	var store TrigramStore
	var text string

	// Store has no trigrams:
	store = NewMapTrigramStore(&TestChooser{})
	text = store.MakeText()

	if text != "" {
		t.Fatalf("Text is invalid. Got %s", text)
	}

	// Store has trigrams:
	store = NewMapTrigramStore(&TestChooser{})

	store.AddTrigram([3]string{"a", "b", "c"})
	store.AddTrigram([3]string{"b", "c", "d"})
	store.AddTrigram([3]string{"b", "c", "e"})

	text = store.MakeText()

	if text != "a b c d" {
		t.Fatalf("Text is invalid. Got %s", text)
	}
}

type TestChooser struct{}

func (c *TestChooser) ChooseInitialTrigram(availableTrigrams TrigramMap) Trigram {
	return Trigram{"a", "b", "c"}
}

func (c *TestChooser) ChooseNextWord(possibleWords map[string]int) string {
	var words []string

	for w := range possibleWords {
		words = append(words, w)
	}

	sort.Strings(words)
	return words[0]
}
