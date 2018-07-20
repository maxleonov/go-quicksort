package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuicksortInt(t *testing.T) {
	const SeqSize = 10

	var sortedInts = make([]int, SeqSize)
	for i := 0; i < SeqSize; i++ {
		sortedInts[i] = i
	}

	var randomInts = make([]interface{}, SeqSize)
	for i, v := range rand.Perm(SeqSize) {
		randomInts[i] = v
	}

	var quickSortedInts = Quicksort(randomInts, func(randomInts, b interface{}) bool {
		return randomInts.(int) < b.(int)
	})

	t.Log(quickSortedInts)

	for i := range quickSortedInts {
		if quickSortedInts[i] != sortedInts[i] {
			t.Errorf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestQuicksortString(t *testing.T) {
	var letters = []string{"K", "R", "A", "T", "E", "L", "E", "P", "U", "I", "M", "Q", "C", "X", "O", "S"}

	var sortedStrings = letters[:]
	sort.Strings(sortedStrings)

	var randomStrings = make([]interface{}, len(letters))
	for i, v := range letters {
		randomStrings[i] = v
	}

	var quickSortedStrings = Quicksort(randomStrings, func(randomStrings, b interface{}) bool {
		return randomStrings.(string) < b.(string)
	})

	t.Log(quickSortedStrings)

	for i := range quickSortedStrings {
		if quickSortedStrings[i] != sortedStrings[i] {
			t.Errorf("Not Sorted!")
			t.Log(quickSortedStrings)
			t.FailNow()
		}
	}
}
