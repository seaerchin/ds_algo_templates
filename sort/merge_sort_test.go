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

// this isn't a rigorous test; should rewrite or have custom test table just for this
func TestMerge(t *testing.T) {
	for _, table := range testTable { 
		mid := len(table.expected) / 2
		left, right := table.expected[:mid], table.expected[mid:]
		got := merge(left, right)
		for i := range got { 
			if got[i] != table.expected[i] { 
				t.Errorf("merge(%v, %v); expected %d, got %d", left, right, table.expected, got)
				break
			}
		}
	}
}
