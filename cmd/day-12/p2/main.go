package main

import (
	"AOC2022/cmd/day-12/djikstra"
	"AOC2022/cmd/day-12/graph"
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	startByte  byte     = 'S'
	endByte    byte     = 'E'
	costNoEdge CostUnit = math.MaxInt
)

type CostUnit int

func gridToNodeNum(x, y, lineLne int) graph.NodeRef {
	return graph.NodeRef(y*lineLne + x)
}

func traverseCost(start, end byte) CostUnit {
	if start == startByte {
		// Current square is the start, this means that it has a height of 'a'
		start = 'a'
	}
	if end == endByte {
		// End square is the end, this means that it has a height of 'z'
		end = 'z'
	}
	if end == startByte {
		// We can't go back to the start
		return costNoEdge
	} else if end <= start+1 {
		// neighbor is at most one higher than c, so we can go there
		return 1
	} else {
		return costNoEdge
	}
}

func main() {

	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	numLines := len(lines)
	lineLen := len(lines[0])
	totalSquares := numLines * lineLen
	pathGraph := graph.New[byte](totalSquares, costNoEdge)
	var end graph.NodeRef
	var starts []graph.NodeRef
	for y := 0; y < numLines; y++ {
		for x := 0; x < lineLen; x++ {
			c := lines[y][x]
			nodeNum := gridToNodeNum(x, y, lineLen)
			if c == startByte {
				// TODO
				starts = append(starts, nodeNum)
			} else if c == endByte {
				// TODO
				end = nodeNum
			} else if c == 'a' {
				starts = append(starts, nodeNum)
			}

			// Check path up
			if y > 0 {
				neighbor := lines[y-1][x]
				neighborNodeNum := gridToNodeNum(x, y-1, lineLen)
				cost := traverseCost(c, neighbor)
				pathGraph.SetEdgeCost(nodeNum, neighborNodeNum, cost, graph.EdgeTypeDirected)
			}
			// Check path down
			if y < numLines-1 {
				neighbor := lines[y+1][x]
				neighborNodeNum := gridToNodeNum(x, y+1, lineLen)
				cost := traverseCost(c, neighbor)
				pathGraph.SetEdgeCost(nodeNum, neighborNodeNum, cost, graph.EdgeTypeDirected)
			}
			// Check path left
			if x > 0 {
				neighbor := lines[y][x-1]
				neighborNodeNum := gridToNodeNum(x-1, y, lineLen)
				cost := traverseCost(c, neighbor)
				pathGraph.SetEdgeCost(nodeNum, neighborNodeNum, cost, graph.EdgeTypeDirected)
			}
			// Check path right
			if x < lineLen-1 {
				neighbor := lines[y][x+1]
				neighborNodeNum := gridToNodeNum(x+1, y, lineLen)
				cost := traverseCost(c, neighbor)
				pathGraph.SetEdgeCost(nodeNum, neighborNodeNum, cost, graph.EdgeTypeDirected)
			}
		}
	}

	fmt.Printf("%d 'a's found\n", len(starts))
	shortest := costNoEdge
	for aNumber, start := range starts {
		dResult, err := djikstra.Djikstra(pathGraph, start, 0, costNoEdge)
		if err != nil {
			panic(err)
		}
		distance := dResult.Distances[end]
		fmt.Printf("Calculating a #%d: %d\n", aNumber, distance)
		if distance != dResult.CostInfinity && distance < shortest {
			shortest = distance
		}
	}
	fmt.Printf("Distance to end: %d\n", shortest)
}
