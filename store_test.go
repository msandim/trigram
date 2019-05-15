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

	/*
		if len(elements) != 1 {
			t.Fatalf("Number of trigrams is not correct")
		}

		elem = elements[[2]string{"a", "b"}]

		if elem.elem != "c" {
			t.Fatalf("Trigram was not added")
		}

		if elem.freq != 1 {
			t.Fatalf("Frequency was not updated")
		}

		store.AddTrigram([3]string{"a", "d", "e"})

		elem = elements[[2]string{"a", "d"}]

		if elem.elem != "e" {
			t.Fatalf("Trigram was not added")
		}

		if elem.freq != 1 {
			t.Fatalf("Frequency was not updated")
		}
	*/
}
