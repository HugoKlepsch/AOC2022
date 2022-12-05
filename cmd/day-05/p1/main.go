package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Crate string

type CrateStack []Crate

func (s *CrateStack) push(c Crate) {
	*s = append(*s, c)
}

func (s *CrateStack) pop() Crate {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *CrateStack) peek() Crate {
	return (*s)[len(*s)-1]
}

type MoveInput struct {
	Source      int
	Destination int
	Times       int
}

func parseCrates(lines []string) map[int]*CrateStack {
	crateStacks := make(map[int]*CrateStack)
	// Get num stacks
	lastLine := lines[len(lines)-1]
	indexes := strings.Split(lastLine, " ")
	lastNum, err := strconv.ParseInt(indexes[len(indexes)-2], 10, 0)
	if err != nil {
		panic(err)
	}
	for i := 0; i < int(lastNum); i++ {
		crateStacks[i] = &CrateStack{}
	}

	// Parse the stacks "vertically"
	stackLines := lines[:len(lines)-1]
	// Reverse the slice so the bottom is the first parse
	for i, j := 0, len(stackLines)-1; i < j; i, j = i+1, j-1 {
		stackLines[i], stackLines[j] = stackLines[j], stackLines[i]
	}

	// Parse each stack vertically
	for stackNum := 0; stackNum < int(lastNum); stackNum++ {
		stackIndex := 1 + (stackNum * 4) // the stacks are separated by 4 characters
		for _, line := range stackLines {
			crate := Crate(fmt.Sprintf("%c", line[stackIndex]))
			if crate != " " {
				crateStacks[stackNum].push(crate)
			}
		}
	}

	return crateStacks
}

func doMove(s map[int]*CrateStack, m MoveInput) {
	for i := 0; i < m.Times; i++ {
		s[m.Destination-1].push(s[m.Source-1].pop())
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

	var crateStacks map[int]*CrateStack

	readingStackInput := true
	stackInput := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if readingStackInput && strings.Contains(line, "[") {
			stackInput = append(stackInput, line)
		} else if readingStackInput {
			// The first line after stacks are done
			readingStackInput = false
			stackInput = append(stackInput, line)
			crateStacks = parseCrates(stackInput)
		} else if line != "" {
			move := MoveInput{}
			_, err := fmt.Sscanf(
				line,
				"move %d from %d to %d",
				&move.Times,
				&move.Source,
				&move.Destination,
			)
			if err != nil {
				panic(err)
			}
			doMove(crateStacks, move)
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	for i := 0; i < len(crateStacks); i++ {
		stack := crateStacks[i]
		fmt.Print(stack.peek())
	}
	fmt.Println()
}
