package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	X      int
	Y      int
	Height int
}

func getViewUp(trees [][]Tree, x, y int) int {
	treeHeight := trees[y][x].Height
	for yi := y - 1; yi >= 0; yi-- {
		if trees[yi][x].Height >= treeHeight {
			return y - yi
		}
	}
	return y
}

func getViewDown(trees [][]Tree, x, y int) int {
	treeHeight := trees[y][x].Height
	gridHeight := len(trees)
	for yi := y + 1; yi < gridHeight; yi++ {
		if trees[yi][x].Height >= treeHeight {
			return yi - y
		}
	}
	return gridHeight - y - 1
}

func getViewLeft(trees [][]Tree, x, y int) int {
	treeHeight := trees[y][x].Height
	for xi := x - 1; xi >= 0; xi-- {
		if trees[y][xi].Height >= treeHeight {
			return x - xi
		}
	}
	return x
}

func getViewRight(trees [][]Tree, x, y int) int {
	treeHeight := trees[y][x].Height
	gridWidth := len(trees[0])
	for xi := x + 1; xi < gridWidth; xi++ {
		if trees[y][xi].Height >= treeHeight {
			return xi - x
		}
	}
	return gridWidth - x - 1
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

	trees := make([][]Tree, 0)
	lineNum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		treeLine := make([]Tree, 0)
		for x, heightRune := range line {
			height, err := strconv.ParseInt(string(heightRune), 10, 0)
			if err != nil {
				panic(err)
			}
			treeLine = append(treeLine, Tree{
				X:      x,
				Y:      lineNum,
				Height: int(height),
			})
		}
		trees = append(trees, treeLine)
		lineNum++
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Calculate visibility for all trees

	gridHeight := len(trees)
	if gridHeight == 0 {
		panic("Invalid grid height")
	}
	gridWidth := len(trees[0])
	if gridWidth == 0 {
		panic("Invalid grid width")
	}

	// view distances
	maxScore := 0
	for y := 1; y < gridHeight-1; y++ {
		for x := 1; x < gridWidth-1; x++ {
			score := getViewUp(trees, x, y) *
				getViewDown(trees, x, y) *
				getViewLeft(trees, x, y) *
				getViewRight(trees, x, y)
			if score > maxScore {
				maxScore = score
			}
			fmt.Printf("(%d, %d): %d\n", x, y, score)
		}
	}
	fmt.Println(maxScore)
}
