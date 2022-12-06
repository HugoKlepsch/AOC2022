package main

import (
	"bufio"
	"fmt"
	"os"
)

type Element rune

type SlidingWindow struct {
	s      []Element
	offset int
}

func (s SlidingWindow) index(i int) int {
	return i % cap(s.s)
}

func (s *SlidingWindow) Enqueue(element Element) {
	if len(s.s) != cap(s.s) {
		s.s = append(s.s, element)
		s.offset++
	} else {
		s.offset = s.index(s.offset)
		s.s[s.offset] = element
		s.offset++
	}
}

func (s *SlidingWindow) Dequeue() Element {
	// TODO
	return Element(0)
}

func (s SlidingWindow) containsOnlyUniqueElements() bool {
	m := make(map[Element]struct{})
	for _, i := range s.s {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = struct{}{}
	}
	return true
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

	for fileScanner.Scan() {
		line := fileScanner.Text()
		slidingWindow := SlidingWindow{
			s: make([]Element, 0, 4),
		}
		for i, e := range line {
			character := string(e)
			_ = character
			slidingWindow.Enqueue(Element(e))
			if i >= 3 && slidingWindow.containsOnlyUniqueElements() {
				println(i + 1)
				return
			}
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
