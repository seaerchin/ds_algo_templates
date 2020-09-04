package sort

import (
	"testing"
)

// needs more test cases; use test table pattern per mergesort
func TestInsertionSort(t *testing.T) {
	arr := []int{2, 3, 6, 1, 3}
	expected := []int{1, 2, 3, 3, 6}
	iter := InsertionSort(arr)
	rec := recursive(arr)
	for idx := range arr {
		if expected[idx] != iter[idx] || expected[idx] != rec[idx] {
			t.Errorf("InsertionSort(%v), expected: %v, iter: %v, rec: %v", arr, expected, iter, rec)
			break
		}
	}
}
