package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	X int
	Y int
}

type Vector Position

func (p Position) vectorAdd(v Vector) Position {
	return Position{
		X: p.X + v.X,
		Y: p.Y + v.Y,
	}
}

func printHeadTail(head, tail Position) {
	var maxX, maxY int
	if head.X > tail.X {
		maxX = head.X + 3
	} else {
		maxX = tail.X + 3
	}
	if head.Y > tail.Y {
		maxY = head.Y + 3
	} else {
		maxY = tail.Y + 3
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if x == head.X && y == head.Y {
				fmt.Printf("H")
			} else if x == tail.X && y == tail.Y {
				fmt.Printf("T")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
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

	head := Position{}
	tail := Position{}

	allTailMoves := make(map[Position]struct{})

	printHeadTail(head, tail)
	fmt.Printf("------\n")
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var direction string
		var moveCount int
		_, err := fmt.Sscanf(line, "%s %d", &direction, &moveCount)
		if err != nil {
			panic(err)
		}

		var moveVector Vector
		if direction == "R" {
			moveVector = Vector{
				X: 1,
				Y: 0,
			}
		} else if direction == "L" {
			moveVector = Vector{
				X: -1,
				Y: 0,
			}
		} else if direction == "U" {
			moveVector = Vector{
				X: 0,
				Y: 1,
			}
		} else if direction == "D" {
			moveVector = Vector{
				X: 0,
				Y: -1,
			}
		}

		fmt.Printf("=== %s %d ===\n", direction, moveCount)
		fmt.Printf("=== (%d,%d) ===\n", moveVector.X, moveVector.Y)
		for i := 0; i < moveCount; i++ {
			head = head.vectorAdd(moveVector)

			printHeadTail(head, tail)

			// Up
			if head.Y > tail.Y+1 {
				tail.Y++
				tail.X = head.X
			}
			// Down
			if head.Y < tail.Y-1 {
				tail.Y--
				tail.X = head.X
			}
			// Right
			if head.X > tail.X+1 {
				tail.X++
				tail.Y = head.Y
			}
			// Left
			if head.X < tail.X-1 {
				tail.X--
				tail.Y = head.Y
			}
			printHeadTail(head, tail)
			fmt.Printf("---\n")

			allTailMoves[tail] = struct{}{}
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(len(allTailMoves))
}
