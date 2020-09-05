package sort

import "testing"

func TestBubbleSort(t *testing.T) { 
	for _, table := range testTable {
		got := BubbleSort(table.input)
		for idx := range got {
			if table.expected[idx] != got[idx] {
				t.Errorf("BubbleSort(%v), expected: %v, got: %v", table.input, table.expected, got)
				break
			}
		}
	}
}