package main

import (
	"testing"
)

func TestAddTrigram(t *testing.T) {

	store := NewTrigramStore()
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
