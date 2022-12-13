package adjacencymatrix

func New(expectedSize int) *AdjacencyMatrix {
	m := AdjacencyMatrix{
		m:        make([][]bool, expectedSize),
		numNodes: expectedSize,
	}
	for i := 0; i < expectedSize; i++ {
		m.m[i] = make([]bool, expectedSize)
	}
	return &m
}

type Node int

type AdjacencyMatrix struct {
	m        [][]bool
	numNodes int
}

// AddNode
// Add a new node to the adjacency matrix. Note that this function is O(n^2),
// where n is the number of nodes in the matrix. If you add each node using this function,
// then that operation is O(n^3). You should specify the expected number of nodes
// when creating the matrix.
// Returns the new node
func (m *AdjacencyMatrix) AddNode() Node {
	newNumNodes := m.numNodes + 1
	m.m = append(m.m, make([]bool, newNumNodes))
	for i := 0; i < m.numNodes; i++ {
		// Resize each existing node line
		m.m[i] = append(m.m[i], false)
	}
	m.numNodes++
	return Node(newNumNodes - 1)
}

func (m *AdjacencyMatrix) SetAdjacentTo(n Node, adjacents ...Node) {
	line := m.m[n]
	for _, adjacent := range adjacents {
		line[adjacent] = true
		m.m[adjacent][n] = true
	}
}

func (m *AdjacencyMatrix) AdjacentTo(n Node) []bool {
	return m.m[n]
}

func (m *AdjacencyMatrix) IsAdjacentTo(n, o Node) bool {
	return m.m[n][o]
}
