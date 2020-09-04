package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	one := []int{1}
	none := []int{}
	scrambled := []int{2, 4, 3, 5, 1}
	testTable := []struct {
		input    []int
		expected []int
	}{
		{sorted, sorted},
		{one, one},
		{none, none},
		{scrambled, sorted},
	}

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
