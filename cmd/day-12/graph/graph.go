package graph

import "math"

const CostNoEdge CostUnit = math.MaxFloat64

type CostUnit float64

func New[T any](expectedSize int) *Graph[T] {
	m := Graph[T]{
		nodes:      make([]T, expectedSize),
		edgeMatrix: make([][]CostUnit, expectedSize),
		NumNodes:   expectedSize,
	}
	for i := 0; i < expectedSize; i++ {
		m.edgeMatrix[i] = make([]CostUnit, expectedSize)
		for j := 0; j < expectedSize; j++ {
			m.edgeMatrix[i][j] = CostNoEdge
		}
	}
	return &m
}

type Graph[T any] struct {
	nodes      []T
	edgeMatrix [][]CostUnit
	NumNodes   int
}

type NodeRef int

func (g *Graph[T]) Get(i NodeRef) T {
	return g.nodes[i]
}

func (g *Graph[T]) Set(i NodeRef, e T) {
	g.nodes[i] = e
}

// AddNode
// Add a new node to the graph. Note that this function is O(n^2),
// where n is the number of nodes in the graph. If you add each node using this function,
// then that operation is O(n^3). You should specify the expected number of nodes
// when creating the graph.
// Returns the new node
func (g *Graph[T]) AddNode(e T) NodeRef {
	newNumNodes := g.NumNodes + 1
	newEdgeLine := make([]CostUnit, newNumNodes)
	for i := 0; i < newNumNodes; i++ {
		newEdgeLine[i] = CostNoEdge
	}
	g.edgeMatrix = append(g.edgeMatrix, newEdgeLine)
	for i := 0; i < g.NumNodes; i++ {
		// Resize each existing node line
		g.edgeMatrix[i] = append(g.edgeMatrix[i], CostNoEdge)
	}
	g.NumNodes++
	return NodeRef(newNumNodes - 1)
}

func (g *Graph[T]) SetEdgeCost(n, other NodeRef, cost CostUnit) {
	g.edgeMatrix[n][other] = cost
	g.edgeMatrix[other][n] = cost
}

func (g *Graph[T]) EdgeCosts(n NodeRef) []CostUnit {
	return g.edgeMatrix[n]
}

func (g *Graph[T]) EdgeCost(n, o NodeRef) CostUnit {
	return g.edgeMatrix[n][o]
}
