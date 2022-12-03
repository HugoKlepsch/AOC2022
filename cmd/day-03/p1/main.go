package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item rune

type Compartment struct {
	Items    []Item
	ItemsMap map[Item]struct{}
}

type Backpack struct {
	Left  Compartment
	Right Compartment
}

func newCompartment() Compartment {
	return Compartment{
		Items:    make([]Item, 0),
		ItemsMap: make(map[Item]struct{}),
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

	score := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		compartmentL := newCompartment()
		for _, c := range line[:len(line)/2] {
			item := Item(c)
			compartmentL.Items = append(compartmentL.Items, item)
			compartmentL.ItemsMap[item] = struct{}{}
		}
		compartmentR := newCompartment()
		for _, c := range line[len(line)/2:] {
			item := Item(c)
			compartmentR.Items = append(compartmentR.Items, item)
			compartmentR.ItemsMap[item] = struct{}{}
		}
		backpack := Backpack{
			Left:  compartmentL,
			Right: compartmentR,
		}
		doubles := backpackDoubles(backpack)

		for _, doubledItem := range doubles {
			score += itemPriority(doubledItem)
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(score)
}

func backpackDoubles(b Backpack) []Item {
	itemDoubles := make([]Item, 0)
	for item, _ := range b.Left.ItemsMap {
		_, ok := b.Right.ItemsMap[item]
		if ok {
			// an item in Left matches an item in Right, this is an error
			itemDoubles = append(itemDoubles, item)
		}
	}
	return itemDoubles
}

func itemPriority(i Item) int {
	if i >= 'A' && i <= 'Z' {
		return int(i) - 38
	}
	if i >= 'a' && i <= 'z' {
		return int(i) - 96
	}
	return 0
}
