package union

// UnionFind consists of two portions - find and union, respectively
// the find portion recursively or iteratively finds the root of a given node
// the union portion unions two distinct sets into a single set
// the outline of union find is given below
// an assumption made here is that the mappings is of the form given below
// and that it maps each node to their immediate parent
func UnionFind(mappings map[int]int) { 
	parents := make([]int, len(mappings))
	for k := range parents { 
		parents[k] = k
	}

	// we use a nested recursive function here but this can be split up as is appropriate
	var find func(cur int) int
	find = func(cur int) int { 
		if parents[cur] != cur {
			parents[cur] = find(parents[cur])
		}
		return parents[cur]
	}

	var union func(x, y int)
	union = func(x, y int) { 
		parents[find(x)] = find(y)
	}

	for k, v := range mappings { 
		union(k, v)
	}
}