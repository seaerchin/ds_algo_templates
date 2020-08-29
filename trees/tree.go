package trees

import (
	"fmt"
	"math"
)

// Node is the data definition for a single tree node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// FromSlice creates a tree from a given int slice;
// a negative value will be assumed to mean that the node itself is empty
func FromSlice(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	return createTree(arr, 0)
}

func createTree(arr []int, idx int) *Node {
	if idx >= len(arr) {
		return nil
	}
	val := arr[idx]
	if val == math.MinInt64 {
		return nil
	}
	node := Node{val, createTree(arr, 2*idx+1), createTree(arr, 2*idx+2)}
	return &node
}

func (t *Node) String() string {
	if t == nil {
		return "nil"
	}
	q := []*Node{t}
	res := ""
	for len(q) > 0 {
		level := len(q)
		for level > 0 {
			cur := q[0]
			q = q[1:]
			if cur.Value == math.MinInt64 {
				res += "nil "
				level--
				continue
			}
			res += fmt.Sprint(cur.Value)
			res += " "
			if cur.Left != nil {
				q = append(q, cur.Left)
			} else {
				q = append(q, &Node{math.MinInt64, nil, nil})
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			} else {
				q = append(q, &Node{math.MinInt64, nil, nil})
			}
			level--
		}
	}
	return res
}
