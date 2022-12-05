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

}

func (s *CrateStack) pop() Crate {
	return Crate("TODO")
}

func parseCrates() {

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

	crateStacks := make(map[int]*CrateStack)

	readingStackInput := true
	stackInput := []string{}
	moveInput := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if readingStackInput && strings.Contains(line, "[") {
			stackInput = append(stackInput, line)
		} else if readingStackInput {
			// The first line after stacks are done
			readingStackInput = false
			indexes := strings.Split(line, " ")
			lastNum, err := strconv.ParseInt(indexes[len(indexes)-1], 10, 0)
			if err != nil {
				panic(err)
			}
			for i := 1; i <= int(lastNum); i++ {
				crateStacks[i] = &CrateStack{}
			}
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
