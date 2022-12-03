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
	Left              Compartment
	Right             Compartment
	ItemsUnion        map[Item]struct{}
	ItemsIntersection map[Item]struct{}
}

func newCompartment() Compartment {
	return Compartment{
		Items:    make([]Item, 0),
		ItemsMap: make(map[Item]struct{}),
	}
}

func lineToBackpack(line string) Backpack {
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
		Left:              compartmentL,
		Right:             compartmentR,
		ItemsUnion:        union(compartmentL.ItemsMap, compartmentR.ItemsMap),
		ItemsIntersection: union(compartmentL.ItemsMap, compartmentR.ItemsMap),
	}
	return backpack
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

		backpack1 := lineToBackpack(line)

		// Just assume that the next two lines exist
		fileScanner.Scan()
		line = fileScanner.Text()
		backpack2 := lineToBackpack(line)

		fileScanner.Scan()
		line = fileScanner.Text()
		backpack3 := lineToBackpack(line)

		i12 := intersection(backpack1.ItemsUnion, backpack2.ItemsUnion)
		fmt.Println("i12: ", i12)
		i3i12 := intersection(i12, backpack3.ItemsUnion)
		fmt.Println("i3i12: ", i3i12)

		if len(i3i12) != 1 {
			err := fmt.Errorf("could not determine badge for group of elves! %v", i3i12)
			panic(err)
		}

		keys := make([]Item, 0, len(i3i12))
		for i := range i3i12 {
			keys = append(keys, i)
		}
		badge := keys[0]
		score += itemPriority(badge)
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(score)
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

func intersection(l, r map[Item]struct{}) map[Item]struct{} {
	intersection := make(map[Item]struct{})
	for item, _ := range l {
		_, ok := r[item]
		if ok {
			// an item in Left matches an item in Right, this is an error
			intersection[item] = struct{}{}
		}
	}
	return intersection
}

func union(l, r map[Item]struct{}) map[Item]struct{} {
	union := make(map[Item]struct{})
	for item, _ := range l {
		union[item] = struct{}{}
	}
	for item, _ := range r {
		union[item] = struct{}{}
	}
	return union
}
