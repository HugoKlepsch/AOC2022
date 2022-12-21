package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ListItem struct {
	l     []ListItem
	num   int
	isNum bool
}

type CompareResult int

const (
	CompareGreater CompareResult = iota
	CompareEqual
	CompareLess
)

func Compare(left, right ListItem) CompareResult {
	if left.isNum && right.isNum {
		// Both integers
		if left.num < right.num {
			return CompareLess
		} else if left.num == right.num {
			return CompareEqual
		} else {
			return CompareGreater
		}
	}

	if !left.isNum && !right.isNum {
		// both lists
		leftLen := len(left.l)
		rightLen := len(right.l)
		i := 0
		for ; i < leftLen && i < rightLen; i++ {
			// i is valid to compare
			compareResult := Compare(left.l[i], right.l[i])
			if compareResult != CompareEqual {
				return compareResult
			}
		}
		if i >= leftLen && i < rightLen {
			// left ran out first
			return CompareLess
		} else if i >= rightLen && i < leftLen {
			// right ran out first
			return CompareGreater
		} else if i >= leftLen && i >= rightLen {
			// They are both the same length, and ran out at the same time
			return CompareEqual
		}
		return CompareEqual
	}

	// One must be a list, and the other must be a number
	var newLeft, newRight ListItem
	if left.isNum && !right.isNum {
		newRight = right
		newLeft = ListItem{
			isNum: false,
			l:     []ListItem{left},
		}
	} else if !left.isNum && right.isNum {
		newLeft = left
		newRight = ListItem{
			isNum: false,
			l:     []ListItem{right},
		}
	} else {
		// This should never happen
		panic("invalid list/not list state")
	}

	return Compare(newLeft, newRight)
}

var justNumberRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]+$`)

func Parse(line string) ListItem {
	list := ListItem{
		l: make([]ListItem, 0),
	}

	// Is this just a number?
	matches := justNumberRegex.FindAllString(line, -1)
	if len(matches) == 1 && matches[0] != "" {
		list.isNum = true
		num, err := strconv.ParseInt(matches[0], 10, 0)
		if err != nil {
			panic(err)
		}
		list.num = int(num)
		return list
	}

	// It must be a list of some type
	parenLevel := 0
	start, end := 0, 0
	for i, r := range line {
		if r == '[' {
			parenLevel++
			if parenLevel == 1 {
				start = i + 1
			}
		} else if r == ']' {
			if parenLevel == 1 && start != end {
				end = i
				list.l = append(list.l, Parse(line[start:end]))
			}
			parenLevel--
		} else if r == ',' && parenLevel == 1 {
			end = i
			list.l = append(list.l, Parse(line[start:end]))
			start = i + 1
		}
	}

	return list
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

	pairNumber := 1
	score := 0

	for fileScanner.Scan() {
		line1 := fileScanner.Text()
		if line1 == "" {
			continue
		}
		list1 := Parse(line1)

		fileScanner.Scan()
		line2 := fileScanner.Text()
		list2 := Parse(line2)

		compareResult := Compare(list1, list2)
		if compareResult == CompareEqual {
			panic("did not determine order")
		} else if compareResult == CompareGreater {
			fmt.Printf("OUT OF ORDER: %s and %s\n", line1, line2)
		} else {
			score += pairNumber
			fmt.Printf("IN ORDER: %s and %s\n", line1, line2)
		}

		pairNumber++
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(score)
}
