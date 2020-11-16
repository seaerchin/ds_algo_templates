package sort

import (
	"testing"
)


func TestCountingSort(t *testing.T) { 
	for _, table := range testTable {
		got := CountingSort(table.input, 10)
		for idx := range got {
			if table.expected[idx] != got[idx] {
				t.Errorf("CountingSort(%v), expected: %v, got: %v", table.input, table.expected, got)
				break
			}
		}
	}
}