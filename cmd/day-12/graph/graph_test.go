package graph

import "testing"

func TestAdjacency(t *testing.T) {
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

	g := New[struct{}](6)

	// Set adjacencies
	g.SetEdgeCost(0, 1, 7)
	g.SetEdgeCost(0, 4, 2)
	g.SetEdgeCost(4, 5, 2)
	g.SetEdgeCost(5, 1, 4)
	g.SetEdgeCost(5, 2, 1)
	g.SetEdgeCost(1, 2, 1)
	g.SetEdgeCost(3, 2, 2)

	// line adjacencies
	assertLineAdjacency([]CostUnit{CostNoEdge, 7, CostNoEdge, CostNoEdge, 2, CostNoEdge}, g.EdgeCosts(0), t)
	assertLineAdjacency([]CostUnit{7, CostNoEdge, 1, CostNoEdge, CostNoEdge, 4}, g.EdgeCosts(1), t)
	assertLineAdjacency([]CostUnit{CostNoEdge, 1, CostNoEdge, 2, CostNoEdge, 1}, g.EdgeCosts(2), t)
	assertLineAdjacency([]CostUnit{CostNoEdge, CostNoEdge, 2, CostNoEdge, CostNoEdge, CostNoEdge}, g.EdgeCosts(3), t)
	assertLineAdjacency([]CostUnit{2, CostNoEdge, CostNoEdge, CostNoEdge, CostNoEdge, 2}, g.EdgeCosts(4), t)
	assertLineAdjacency([]CostUnit{CostNoEdge, 4, 1, CostNoEdge, 2, CostNoEdge}, g.EdgeCosts(5), t)

	// Exact adjacency
	type TestCase struct {
		node, other NodeRef
		cost        CostUnit
	}
	testCases := []TestCase{
		{
			node:  0,
			other: 1,
			cost:  7,
		},
		{
			node:  0,
			other: 4,
			cost:  2,
		},
		{
			node:  0,
			other: 5,
			cost:  CostNoEdge,
		},
		{
			node:  4,
			other: 0,
			cost:  2,
		},
		{
			node:  4,
			other: 1,
			cost:  CostNoEdge,
		},
	}
	for _, testCase := range testCases {
		actual := g.EdgeCost(testCase.node, testCase.other)
		if testCase.cost != actual {
			t.Errorf("EdgeCost(%d, %d) -> %f, expected %f", testCase.node, testCase.other, actual, testCase.cost)
		}
	}
}

func assertLineAdjacency(expected, actual []CostUnit, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("invalid number of edges")
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("cost to (%d): %f, expected %f\n", i, actual[i], expected[i])
		}
	}
}

func TestSetGet(t *testing.T) {
	g := New[float64](6)

	// Set values
	for i := 0; i < 6; i++ {
		g.Set(NodeRef(i), float64(i+10))
	}

	// Get values
	for i := 0; i < 6; i++ {
		actual := g.Get(NodeRef(i))
		expected := float64(i + 10)
		if actual != expected {
			t.Errorf("get (%d): %f, expected %f\n", i, actual, expected)
		}
	}
}
