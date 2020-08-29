package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 3, 6, 1, 3}
	expected := []int{1, 2, 3, 3, 6}
	got := InsertionSort(arr)
	for idx := range arr {
		if expected[idx] != got[idx] {
			t.Errorf("InsertionSort(%v), expected: %v, got: %v", arr, expected, got)
			break
		}
	}
}
