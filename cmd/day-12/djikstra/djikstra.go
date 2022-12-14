package djikstra

import (
	"AOC2022/cmd/day-12/graph"
	"AOC2022/cmd/day-12/priorityqueue"

	"golang.org/x/exp/constraints"
)

const (
	NodeRefUndefined graph.NodeRef = -1
)

type DjikstraResult[CostUnit constraints.Ordered] struct {
	CostInfinity CostUnit
	Distances    []CostUnit
	Previous     []graph.NodeRef
}

// Djikstra
// Based off of pseudocode on Wikipedia:
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Using_a_priority_queue
func Djikstra[T any, CostUnit constraints.Ordered](
	g *graph.Graph[T, CostUnit],
	source graph.NodeRef,
	zeroCost, infiniteCost CostUnit,
) (*DjikstraResult[CostUnit], error) {
	numNodes := g.NumNodes

	d := DjikstraResult[CostUnit]{
		CostInfinity: infiniteCost,
		Distances:    make([]CostUnit, numNodes),
		Previous:     make([]graph.NodeRef, numNodes),
	}

	// "min" priority queue, for grabbing the lowest distance
	nodePriorityQ := priorityqueue.NewHeapPriorityQueue[graph.NodeRef, CostUnit](false)

	d.Distances[source] = zeroCost
	d.Previous[source] = NodeRefUndefined

	for node := graph.NodeRef(0); node < graph.NodeRef(numNodes); node++ {
		if node != source {
			d.Distances[node] = infiniteCost    // Unknown distance from source -> node
			d.Previous[node] = NodeRefUndefined // Unknown predecessor from source -> node
		}
		nodePriorityQ.Insert(node, d.Distances[node])
	}

	for !nodePriorityQ.IsEmpty() {
		node, _, err := nodePriorityQ.Pop()
		if err != nil {
			return nil, err
		}
		edgeCosts := g.EdgeCosts(node)
		for neighbor, edgeCost := range edgeCosts {
			if edgeCost != g.CostNoEdge {
				// This node is a neighbor of the current node pulled from the priority Q

				// Calculate the distance to the neighbor via this node
				alternatePathCost := d.Distances[node] + edgeCost

				if alternatePathCost < d.Distances[neighbor] {
					// We found a shorter path to the neighbor! Update the neighbor with the way back
					d.Distances[neighbor] = alternatePathCost
					d.Previous[neighbor] = node
					nodePriorityQ.SetPriority(graph.NodeRef(neighbor), alternatePathCost)
				}
			}
		}
	}
	return &d, nil
}

func PathFromPrevious(n graph.NodeRef, previous []graph.NodeRef) []graph.NodeRef {
	nodePath := make([]graph.NodeRef, 0)

	// Traverse the path back from n to the source
	var pathNode graph.NodeRef = n
	for pathNode != NodeRefUndefined {
		nodePath = append(nodePath, pathNode)
		pathNode = previous[pathNode]
	}

	// The path is reversed, so lets unreverse it
	pathLen := len(nodePath)
	for i := 0; i < (pathLen / 2); i++ {
		tmp := nodePath[i]
		nodePath[i] = nodePath[pathLen-i-1]
		nodePath[pathLen-i-1] = tmp
	}
	return nodePath
}
