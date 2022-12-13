package adjacencymatrix

import "testing"

func TestAM(t *testing.T) {
	// Matrix:
	//      0
	//    /   \
	//   4     1
	//    \   /|
	//      5  |
	//        \|
	//     3---2

	m := New(6)
	m.SetAdjacentTo(0, 1, 4)
	m.SetAdjacentTo(1, 0, 2, 5)
	m.SetAdjacentTo(2, 1, 3, 5)
	m.SetAdjacentTo(3, 2)
	m.SetAdjacentTo(4, 0, 5)
	m.SetAdjacentTo(5, 1, 2, 4)

	// line adjacencies
	assertLineAdjacency([]bool{false, true, false, false, true, false}, m.AdjacentTo(0), t)
	assertLineAdjacency([]bool{true, false, true, false, false, true}, m.AdjacentTo(1), t)
	assertLineAdjacency([]bool{false, true, false, true, false, true}, m.AdjacentTo(2), t)
	assertLineAdjacency([]bool{false, false, true, false, false, false}, m.AdjacentTo(3), t)
	assertLineAdjacency([]bool{true, false, false, false, false, true}, m.AdjacentTo(4), t)
	assertLineAdjacency([]bool{false, true, true, false, true, false}, m.AdjacentTo(5), t)

	// Exact adjacency
	type TestCase struct {
		node, other Node
		adjacent    bool
	}
	testCases := []TestCase{
		{
			node:     0,
			other:    1,
			adjacent: true,
		},
		{
			node:     0,
			other:    4,
			adjacent: true,
		},
		{
			node:     0,
			other:    5,
			adjacent: false,
		},
		{
			node:     4,
			other:    0,
			adjacent: true,
		},
		{
			node:     4,
			other:    1,
			adjacent: false,
		},
	}
	for _, testCase := range testCases {
		actual := m.IsAdjacentTo(testCase.node, testCase.other)
		if testCase.adjacent != actual {
			t.Fatalf("IsAdjacentTo(%d, %d) -> %t, expected %t", testCase.node, testCase.other, actual, testCase.adjacent)
		}
	}
}

func assertLineAdjacency(expected, actual []bool, t *testing.T) {
	if len(expected) != len(actual) {
		t.Fail()
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Fail()
		}
	}
}
