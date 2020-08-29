package search

import (
	tree "ds_algo_templates/trees"
	"testing"
)

func TestDFS(t *testing.T) {
	nilRoot := &tree.Node{}
	testTable := []struct {
		root     *tree.Node
		target   int
		expected bool
	}{
		{nilRoot, 1, false},
	}

	for _, table := range testTable {
		output := DFS(table.root, table.target)
		if output != table.expected {
			t.Errorf("DFS(%v, %d), expected: %t, got: %t", table.root, table.target, table.expected, output)
		}
	}
}
