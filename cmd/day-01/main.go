package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	calorieMap := make(map[int]int)

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elfIndex := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elfIndex++
			continue
		}
		calories64, err := strconv.ParseInt(line, 10, 0)
		calories := int(calories64)

		if err != nil {
			panic(err)
		}

		existingCalories, ok := calorieMap[elfIndex]
		if !ok {
			existingCalories = 0
		}
		calorieMap[elfIndex] = existingCalories + calories
	}

	calorieSlice := make([]int, 0)
	for _, v := range calorieMap {
		calorieSlice = append(calorieSlice, v)
	}

	sort.Ints(calorieSlice)

	for _, i := range calorieSlice {
		fmt.Println(i)
	}

	readFile.Close()
}
