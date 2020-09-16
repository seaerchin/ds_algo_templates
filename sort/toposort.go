package sort

import graph "ds_algo_templates/data_structures/graph"

// TODO: test this
func Toposort(g graph.Graph) []int {
	n := len(g.Nodes())
	seen := make(map[int]bool)
	ordering := make([]int, n)

	for at := 0; at < n; at++ {
		if !seen[at] {
			visited := []int{}

		}
	}
}

func dfs(g graph.Graph, seen map[int]bool, results []int, at graph.Node) {

}
