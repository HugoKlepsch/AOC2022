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
	// (0, 1) = 6
	// (0, 4) = 2
	// (4, 5) = 2
	// (5, 1) = 4
	// (5, 2) = 1
	// (1, 2) = 1
	// (3, 2) = 2

	m := New[struct{}](6)

	// Set adjacencies
	m.SetEdgeCost(0, 1, 6)
	m.SetEdgeCost(0, 4, 2)
	m.SetEdgeCost(4, 5, 2)
	m.SetEdgeCost(5, 1, 4)
	m.SetEdgeCost(5, 2, 1)
	m.SetEdgeCost(1, 2, 1)
	m.SetEdgeCost(3, 2, 2)

	// line adjacencies
	assertLineAdjacency([]CostUnit{CostNoEdge, 6, CostNoEdge, CostNoEdge, 2, CostNoEdge}, m.EdgeCosts(0), t)
	assertLineAdjacency([]CostUnit{6, CostNoEdge, 1, CostNoEdge, CostNoEdge, 4}, m.EdgeCosts(1), t)
	assertLineAdjacency([]CostUnit{CostNoEdge, 1, CostNoEdge, 2, CostNoEdge, 1}, m.EdgeCosts(2), t)
	assertLineAdjacency([]CostUnit{CostNoEdge, CostNoEdge, 2, CostNoEdge, CostNoEdge, CostNoEdge}, m.EdgeCosts(3), t)
	assertLineAdjacency([]CostUnit{2, CostNoEdge, CostNoEdge, CostNoEdge, CostNoEdge, 2}, m.EdgeCosts(4), t)
	assertLineAdjacency([]CostUnit{CostNoEdge, 4, 1, CostNoEdge, 2, CostNoEdge}, m.EdgeCosts(5), t)

	// Exact adjacency
	type TestCase struct {
		node, other NodeRef
		cost        CostUnit
	}
	testCases := []TestCase{
		{
			node:  0,
			other: 1,
			cost:  6,
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
		actual := m.EdgeCost(testCase.node, testCase.other)
		if testCase.cost != actual {
			t.Fatalf("EdgeCost(%d, %d) -> %f, expected %f", testCase.node, testCase.other, actual, testCase.cost)
		}
	}
}

func assertLineAdjacency(expected, actual []CostUnit, t *testing.T) {
	if len(expected) != len(actual) {
		t.Fail()
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Fail()
		}
	}
}

func TestSetGet(t *testing.T) {
	m := New[float64](6)

	// Set values
	for i := 0; i < 6; i++ {
		m.Set(NodeRef(i), float64(i+10))
	}

	// Get values
	for i := 0; i < 6; i++ {
		actual := m.Get(NodeRef(i))
		if actual != float64(i+10) {
			t.Fail()
		}
	}
}
