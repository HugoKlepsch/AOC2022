package main

import (
	"AOC2022/cmd/day-11/monkey"
	"AOC2022/cmd/day-11/ringbuffer"
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

func main() {

	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	monkeys := make([]*monkey.Monkey, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		lines := []string{line}
		for i := 0; i < 6; i++ {
			fileScanner.Scan()
			line = fileScanner.Text()
			lines = append(lines, line)
		}
		monkey, err := monkey.ParseLinesToMonkey(lines)
		if err != nil {
			panic(err)
		}
		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 20; i++ {
		for monkeyNum, monkey := range monkeys {
			fmt.Printf("Monkey %d:\n", monkeyNum)
			for item, err := monkey.Items.Dequeue(); err == nil; item, err = monkey.Items.Dequeue() {
				fmt.Printf("\tMonkey inspects an item with a worry level of %d.\n", item)
				newItem := monkey.Operation.Do(item)
				fmt.Printf("\t\tWorry level is changed to %d.\n", newItem)
				newItem /= 3
				fmt.Printf("\t\tWorry level relaxes to %d.\n", newItem)
				throwTarget := monkey.Test.Route(newItem)
				fmt.Printf("\t\tItem with worry %d routed to monkey %d.\n", newItem, throwTarget)
				err2 := monkeys[throwTarget].Items.Enqueue(newItem)
				if err2 != nil {
					panic(err2)
				}
				monkey.NItemsInspected++
			}
			if err != nil && !errors.Is(err, ringbuffer.ErrBufferEmpty) {
				panic(err)
			}
		}
	}

	for monkeyNum, monkey := range monkeys {
		fmt.Printf("Monkey: %d, NItems: %d\n", monkeyNum, monkey.NItemsInspected)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].NItemsInspected < monkeys[j].NItemsInspected
	})
	a := monkeys[len(monkeys)-2].NItemsInspected
	b := monkeys[len(monkeys)-1].NItemsInspected
	fmt.Printf("Top two: %d * %d = %d\n", a, b, a*b)

	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
