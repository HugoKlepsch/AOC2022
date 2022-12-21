package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

	packets := make([]ListItem, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		list := Parse(line)
		packets = append(packets, list)
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	twoDivider := ListItem{
		l: []ListItem{
			{
				l: []ListItem{
					{
						num: 2,
						isNum: true,
					},
				},
				isNum: false,
			},
		},
		isNum: false,
	}
	sixDivider := ListItem{
		l: []ListItem{
			{
				l: []ListItem{
					{
						num: 6,
						isNum: true,
					},
				},
				isNum: false,
			},
		},
		isNum: false,
	}

	packets = append(packets, twoDivider)
	packets = append(packets, sixDivider)

	sort.Slice(packets, func(i, j int) bool {
		compareResult := Compare(packets[i], packets[j])
		return compareResult == CompareLess
	})

	twoIndex, sixIndex := 0, 0
	for i := 0; i < len(packets); i++ {
		if Compare(packets[i], twoDivider) == CompareEqual {
			twoIndex = i + 1 // these are "one based"
		} else if Compare(packets[i], sixDivider) == CompareEqual {
			sixIndex = i + 1 // these are "one based"
		}
	}

	fmt.Println(twoIndex * sixIndex)
}
