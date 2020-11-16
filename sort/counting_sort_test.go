package sort

import (
	"testing"
)


func TestCountingSort(t *testing.T) { 
	for _, table := range testTable {
		got := CountingSort(table.input, 10) // instead of 10, this should be the largest element in the array
		for idx := range got {
			if table.expected[idx] != got[idx] {
				t.Errorf("CountingSort(%v), expected: %v, got: %v", table.input, table.expected, got)
				break
			}
		}
	}
}