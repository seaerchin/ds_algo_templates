package search

import "testing"

func TestBinarySearch(t *testing.T) {
	ordered := []int{0, 1, 2}
	empty := []int{}
	single := []int{0}

	testTable := []struct {
		arr      []int
		target   int
		expected int
	}{
		// happy path
		{ordered, 0, 0},
		// not found  in array
		{ordered, 3, -1},
		// check on empty arr
		{empty, 0, -1},
		// checks for single element arr
		{single, 0, 0},
		{single, 1, -1},
	}

	for _, table := range testTable {
		testIdx := BinarySearch(table.arr, table.target)
		if testIdx != table.expected {
			t.Errorf("BinarySearch(%v, %d), expected: %d, got: %d", table.arr, table.target, table.expected, testIdx)
		}
	}
}
