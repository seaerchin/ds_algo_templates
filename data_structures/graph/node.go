package graph

// credits to https://github.com/gonum/gonum/blob/master/graph/graph.go
// inspiration for the implementation shown here
// TODO: add impl for weighted graphs

type Edge interface {
	From() Node
	To() Node
}

// Node returns an ID to a external struct that contains our associated data/meta-data
// callers define own struct with payload to utilize this
type Node interface {
	ID() int
}

// WeightedEdge is also, by definition, a directed edge
type WeightedEdge interface {
	Edge
	Weight() int
}

// TODO: improve this definition for a graph
type Graph interface {
	Nodes() []Node
	EdgesOf(n Node) map[Node]Edge
}

type unweightedGraph struct {
	nodes map[int]Node
	// given a node id, return a mapping from neighbour nodes to their edge
	edges map[int]map[Node]Edge
}

type simpleNode int

type simpleEdge struct {
	to   Node
	from Node
}

func (e *simpleEdge) To() Node {
	return e.to
}

func (e *simpleEdge) From() Node {
	return e.from
}

func (s simpleNode) ID() int {
	return int(s)
}

func (ug *unweightedGraph) Nodes() []Node {
	nodes := make([]Node, 0)
	for _, v := range ug.nodes {
		nodes = append(nodes, v)
	}
	return nodes
}

func (ug *unweightedGraph) EdgesOf(n Node) map[Node]Edge {
	return ug.edges[n.ID()]
}

// keep unweightedGraph hidden and expose this to callers
func makeUnweighted(adjMatrix map[int]int) Graph {
	nodes := make(map[int]Node)
	edges := make(map[int]map[Node]Edge)
	for to, from := range adjMatrix {
		nodes[to] = simpleNode(to)
		edges[to][simpleNode(from)] = &simpleEdge{simpleNode(to), simpleNode(from)}
	}
	return &unweightedGraph{
		nodes,
		edges,
	}
}
