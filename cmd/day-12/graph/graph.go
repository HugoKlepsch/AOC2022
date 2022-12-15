package graph

import (
	"golang.org/x/exp/constraints"
)

type EdgeType int

const (
	EdgeTypeDirected EdgeType = iota
	EdgeTypeBiDirectional
)

func New[T any, CostUnit constraints.Ordered](expectedSize int, costNoEdge CostUnit) *Graph[T, CostUnit] {
	g := Graph[T, CostUnit]{
		nodes:      make([]T, expectedSize),
		edgeMatrix: make([][]CostUnit, expectedSize),
		NumNodes:   expectedSize,
		CostNoEdge: costNoEdge,
	}
	for i := 0; i < expectedSize; i++ {
		g.edgeMatrix[i] = make([]CostUnit, expectedSize)
		for j := 0; j < expectedSize; j++ {
			g.edgeMatrix[i][j] = g.CostNoEdge
		}
	}
	return &g
}

type Graph[T any, CostUnit constraints.Ordered] struct {
	nodes      []T
	edgeMatrix [][]CostUnit
	NumNodes   int
	CostNoEdge CostUnit
}

type NodeRef int

func (g *Graph[T, CostUnit]) Get(i NodeRef) T {
	return g.nodes[i]
}

func (g *Graph[T, CostUnit]) Set(i NodeRef, e T) {
	g.nodes[i] = e
}

// AddNode
// Add a new node to the graph. Note that this function is O(n^2),
// where n is the number of nodes in the graph. If you add each node using this function,
// then that operation is O(n^3). You should specify the expected number of nodes
// when creating the graph.
// Returns the new node
func (g *Graph[T, CostUnit]) AddNode(e T) NodeRef {
	newNumNodes := g.NumNodes + 1
	newEdgeLine := make([]CostUnit, newNumNodes)
	for i := 0; i < newNumNodes; i++ {
		newEdgeLine[i] = g.CostNoEdge
	}
	g.edgeMatrix = append(g.edgeMatrix, newEdgeLine)
	for i := 0; i < g.NumNodes; i++ {
		// Resize each existing node line
		g.edgeMatrix[i] = append(g.edgeMatrix[i], g.CostNoEdge)
	}
	g.NumNodes++
	return NodeRef(newNumNodes - 1)
}

func (g *Graph[T, CostUnit]) SetEdgeCost(n, other NodeRef, cost CostUnit, edgeType EdgeType) {
	g.edgeMatrix[n][other] = cost
	if edgeType == EdgeTypeBiDirectional {
		g.edgeMatrix[other][n] = cost
	}
}

func (g *Graph[T, CostUnit]) EdgeCosts(n NodeRef) []CostUnit {
	return g.edgeMatrix[n]
}

func (g *Graph[T, CostUnit]) EdgeCost(n, o NodeRef) CostUnit {
	return g.edgeMatrix[n][o]
}
