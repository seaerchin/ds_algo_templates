package sort

import (
	"testing"
)

// needs more test cases; use test table pattern per mergesort
func TestInsertionSort(t *testing.T) {
	for _, testEntry := range testTable {
		iter := InsertionSort(testEntry.input)
		rec := recursive(testEntry.input)	
		for idx := range iter {
			if testEntry.expected[idx] != iter[idx] || testEntry.expected[idx] != rec[idx] {
				t.Errorf("InsertionSort(%v), expected: %v, iter: %v, rec: %v", testEntry.input, testEntry.expected, iter, rec)
				break
			}
		}
	}
}