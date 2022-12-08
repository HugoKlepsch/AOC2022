package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	X        int
	Y        int
	Height   int
	VisTop   bool
	VisBot   bool
	VisLeft  bool
	VisRight bool
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
				X:        x,
				Y:        lineNum,
				Height:   int(height),
				VisTop:   false,
				VisBot:   false,
				VisLeft:  false,
				VisRight: false,
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

	// Visibililty from the Top
	for x := 0; x < gridWidth; x++ {
		// X counts each column
		// Y counts each row in that column

		highestTree := trees[0][x].Height
		trees[0][x].VisTop = true
		// Start at 1 because the first row is always visible
		for y := 1; y < gridHeight; y++ {
			currentTree := trees[y][x]
			if currentTree.Height > highestTree {
				highestTree = currentTree.Height
				currentTree.VisTop = true
				trees[y][x] = currentTree
			}
		}
	}

	// Visibililty from the Left
	for y := 0; y < gridHeight; y++ {
		// Y counts each row
		// X counts each column in that row

		highestTree := trees[y][0].Height
		trees[y][0].VisLeft = true
		// Start at 1 because the first col is always visible
		for x := 1; x < gridWidth; x++ {
			currentTree := trees[y][x]
			if currentTree.Height > highestTree {
				highestTree = currentTree.Height
				currentTree.VisLeft = true
				trees[y][x] = currentTree
			}
		}
	}

	// Visibililty from the Right
	for y := 0; y < gridHeight; y++ {
		// Y counts each row
		// X counts each column in that row

		highestTree := trees[y][gridWidth-1].Height
		trees[y][gridWidth-1].VisLeft = true
		// Start at -2 because the first col is always visible
		for x := gridWidth - 2; x >= 0; x-- {
			currentTree := trees[y][x]
			if currentTree.Height > highestTree {
				highestTree = currentTree.Height
				currentTree.VisRight = true
				trees[y][x] = currentTree
			}
		}
	}

	// Visibililty from the Bot
	for x := 0; x < gridWidth; x++ {
		// X counts each column
		// Y counts each row in that column

		highestTree := trees[gridHeight-1][x].Height
		trees[gridHeight-1][x].VisTop = true
		// Start at -2 because the first row is always visible
		for y := gridHeight - 2; y >= 0; y-- {
			currentTree := trees[y][x]
			if currentTree.Height > highestTree {
				highestTree = currentTree.Height
				currentTree.VisBot = true
				trees[y][x] = currentTree
			}
		}
	}

	countVisible := 0
	// Count visible
	for _, treeLine := range trees {
		for _, tree := range treeLine {
			if tree.VisTop || tree.VisLeft || tree.VisBot || tree.VisRight {
				countVisible++
			}
		}
	}

	fmt.Println(countVisible)
}
