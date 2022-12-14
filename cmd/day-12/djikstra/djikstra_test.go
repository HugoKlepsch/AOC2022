package djikstra

import (
	"AOC2022/cmd/day-12/graph"
	"math"
	"testing"

	"golang.org/x/exp/constraints"
)

func testGraph() *graph.Graph[struct{}, int8] {
	// Matrix:
	//      0
	//    /   \
	//   4     1
	//    \   /|
	//      5  |
	//        \|
	//     3---2
	//
	// Costs:
	// (0, 1) = 7
	// (0, 4) = 2
	// (4, 5) = 2
	// (5, 1) = 4
	// (5, 2) = 1
	// (1, 2) = 1
	// (3, 2) = 2

	g := graph.New[struct{}, int8](6, math.MaxInt8)

	// Set adjacencies
	g.SetEdgeCost(0, 1, 7)
	g.SetEdgeCost(0, 4, 2)
	g.SetEdgeCost(4, 5, 2)
	g.SetEdgeCost(5, 1, 4)
	g.SetEdgeCost(5, 2, 1)
	g.SetEdgeCost(1, 2, 1)
	g.SetEdgeCost(3, 2, 2)

	return g
}

func assertDistance[CostUnit constraints.Ordered](node graph.NodeRef, expected CostUnit,
	distances []CostUnit, t *testing.T) {
	actual := distances[node]
	if expected != actual {
		t.Errorf("Distance to (%d): %v, expected %v\n", node, actual, expected)
	}
}

func assertPath(node graph.NodeRef, expected []graph.NodeRef,
	previous []graph.NodeRef, t *testing.T) {
	if node != expected[len(expected)-1] {
		t.Fatal("node must be the last element of the expected path")
	}
	nodePath := PathFromPrevious(node, previous)
	pathLen := len(nodePath)
	expectedLen := len(expected)
	if pathLen != expectedLen {
		t.Fatalf("len(nodePath): %d != len(expected): %d\n", pathLen, expectedLen)
	}
	for i := 0; i < pathLen; i++ {
		expectedNode := expected[i]
		actualNode := nodePath[i]
		if expectedNode != actualNode {
			t.Fatalf("Step %d expected node %d, got %d\n", i, expectedNode, actualNode)
		}
	}
}

func TestDjikstra(t *testing.T) {
	g := testGraph()

	var source graph.NodeRef = 0
	djikstraResult, err := Djikstra(g, source, 0, math.MaxInt8)
	if err != nil {
		t.Error(err)
	}

	assertDistance(0, 0, djikstraResult.Distances, t)
	assertDistance(1, 6, djikstraResult.Distances, t)
	assertDistance(2, 5, djikstraResult.Distances, t)
	assertDistance(3, 7, djikstraResult.Distances, t)
	assertDistance(4, 2, djikstraResult.Distances, t)
	assertDistance(5, 4, djikstraResult.Distances, t)

	assertPath(0, []graph.NodeRef{0}, djikstraResult.Previous, t)
	assertPath(1, []graph.NodeRef{0, 4, 5, 2, 1}, djikstraResult.Previous, t)
	assertPath(2, []graph.NodeRef{0, 4, 5, 2}, djikstraResult.Previous, t)
	assertPath(3, []graph.NodeRef{0, 4, 5, 2, 3}, djikstraResult.Previous, t)
	assertPath(4, []graph.NodeRef{0, 4}, djikstraResult.Previous, t)
	assertPath(5, []graph.NodeRef{0, 4, 5}, djikstraResult.Previous, t)
}
