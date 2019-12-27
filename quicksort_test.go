package main

import "testing"

func TestQuicksort(t *testing.T) {
	expected := [5]int{1, 2, 3, 4, 5}
	testvalues := [5]int{2, 5, 3, 1, 4}

	QuickSort(testvalues[:])
	if testvalues == expected {
		t.Errorf("Array %v is not sorted!", testvalues)
	}
}
