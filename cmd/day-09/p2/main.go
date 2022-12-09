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

func printHeadTail(head Position, tails []Position) {
	var maxX, maxY int

	maxX = head.X
	maxY = head.Y
	for _, tail := range tails {
		if tail.X > maxX {
			maxX = tail.X
		}
		if tail.Y > maxY {
			maxY = tail.Y
		}
	}
	maxX += 3
	maxY += 3

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if x == head.X && y == head.Y {
				fmt.Printf("H")
			} else {
				var found = false
				for i, tail := range tails {
					if x == tail.X && y == tail.Y {
						fmt.Printf("%d", i+1)
						found = true
						break
					}
				}
				if !found {
					fmt.Printf(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveRope(head, tail Position) Position {
	// Up
	if head.Y > tail.Y+1 {
		tail.Y = head.Y - 1
		if head.X > tail.X+1 {
			tail.X = head.X - 1
		} else if head.X < tail.X-1 {
			tail.X = head.X + 1
		} else {
			tail.X = head.X
		}
	}
	// Down
	if head.Y < tail.Y-1 {
		tail.Y = head.Y + 1
		if head.X > tail.X+1 {
			tail.X = head.X - 1
		} else if head.X < tail.X-1 {
			tail.X = head.X + 1
		} else {
			tail.X = head.X
		}
	}
	// Right
	if head.X > tail.X+1 {
		tail.X = head.X - 1
		if head.Y > tail.Y+1 {
			tail.Y = head.Y - 1
		} else if head.Y < tail.Y-1 {
			tail.Y = head.Y + 1
		} else {
			tail.Y = head.Y
		}
	}
	// Left
	if head.X < tail.X-1 {
		tail.X = head.X + 1
		if head.Y > tail.Y+1 {
			tail.Y = head.Y - 1
		} else if head.Y < tail.Y-1 {
			tail.Y = head.Y + 1
		} else {
			tail.Y = head.Y
		}
	}
	return tail
}

func TestMoveRopeDiagonally() {
	head := Position{
		X: 2,
		Y: 2,
	}
	tail := Position{}

	tail = moveRope(head, tail)
}

func main() {
	// TestMoveRopeDiagonally()

	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	head := Position{}
	one := Position{}
	two := Position{}
	three := Position{}
	four := Position{}
	five := Position{}
	six := Position{}
	seven := Position{}
	eight := Position{}
	tail := Position{}

	allTailMoves := make(map[Position]struct{})

	printHeadTail(head, []Position{
		one, two, three, four, five, six, seven, eight, tail,
	})
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
			printHeadTail(head, []Position{
				one, two, three, four, five, six, seven, eight, tail,
			})

			one = moveRope(head, one)
			two = moveRope(one, two)
			three = moveRope(two, three)
			four = moveRope(three, four)
			five = moveRope(four, five)
			six = moveRope(five, six)
			seven = moveRope(six, seven)
			eight = moveRope(seven, eight)
			tail = moveRope(eight, tail)
			printHeadTail(head, []Position{
				one, two, three, four, five, six, seven, eight, tail,
			})
			fmt.Printf("---\n")

			allTailMoves[tail] = struct{}{}
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// between 2557 and 2617
	fmt.Println(len(allTailMoves))
}
