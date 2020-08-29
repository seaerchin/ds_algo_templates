package search

import tree "ds_algo_templates/trees"

// BFS (breadth first search) uses traverses a binary tree and does some work on each node;
// unlike DFS which expands the immediate node until it hits the base case (in this case when node is nil)
// BFS expands the set of children nodes for all nodes on the current level
// the runtime of BFS is O(nodes + edges)
func BFS(root *tree.Node, target int) bool {
	if root == nil {
		return false
	}
	// BFS uses a queue to store the current level's nodes
	q := []*tree.Node{root}
	for len(q) > 0 {
		// number of items on current level
		level := len(q)
		for level > 0 {
			cur, q := q[0], q[1:]
			if cur.Value == target {
				return true
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return false
}
