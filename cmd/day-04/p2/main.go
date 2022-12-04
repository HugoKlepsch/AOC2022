package main

import (
	"bufio"
	"fmt"
	"os"
)

type SectionRange struct {
	Start int
	End   int
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

	score := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		rangeA := SectionRange{}
		rangeB := SectionRange{}
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &rangeA.Start, &rangeA.End, &rangeB.Start, &rangeB.End)
		if err != nil {
			panic(err)
		}

		if rangeOverlaps(rangeA, rangeB) {
			score += 1
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(score)
}

func rangeOverlaps(a, b SectionRange) bool {
	return a.End >= b.Start && a.Start <= b.End
}
