package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, table := range testTable {
		got := MergeSort(table.input)
		for idx := range got {
			if table.expected[idx] != got[idx] {
				t.Errorf("MergeSort(%v), expected: %v, got: %v", got, table.expected, got)
				break
			}
		}
	}
}

// this isn't a rigorous test; should rewrite
func TestMerge(t *testing.T) {
	left, right := []int{1, 4}, []int{2, 2, 3}
	got := merge(left, right)
	expected := []int{1, 2, 2, 3, 4}
	for k := range expected {
		if got[k] != expected[k] {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}
}
