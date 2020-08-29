package search

import tree "ds_algo_templates/trees"

// DFS (depth first search) uses traverses a binary tree and does some work on each node;
// unlike BFS which expands the set of immediate nodes, DFS proceeds by recursion of the immediate left/right node
// the runtime of DFS is O(nodes + edges)
func DFS(root *tree.Node, target int) bool {
	if root == nil {
		return false
	}
	// check then recur on children of our current node
	// order of traversal is pre-order; we could do dfs in-order also
	// this could be done through recurring on left then check value of root then recur on right
	if root.Value == target {
		return true
	}
	if DFS(root.Left, target) || DFS(root.Right, target) {
		return true
	}
	return false
}
