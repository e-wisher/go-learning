package main

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := range a {
		if j != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	array := []int{2, 23, 8, 4, 10, 12}
	want := []int{2, 4, 8, 10, 12, 23}
	sorted, err := BubbleSort(array)
	checkslices := equal(sorted, want)
	if err != nil || !checkslices == true {
		t.Fatalf(`BubbleSort(array) = %q, want match for %v`, sorted, want)
	}
}

func TestEmptyBubbleSort(t *testing.T) {
	array := []int{}
	sorted, err := BubbleSort(array)
	checkslices := equal(sorted, array)
	if !checkslices == true || err == nil {
		t.Fatalf(`BubbleSort(array) = %q, want match for []`, array)
	}
}
