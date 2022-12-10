package main

import (
	"AOC2022/cmd/day-10/engine"
	"AOC2022/cmd/day-10/engine/opcodes"
	"bufio"
	"fmt"
	"os"
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

	e := engine.NewEngine()
	e.AddOpCode(&opcodes.OpCodeAddX{})
	e.AddOpCode(&opcodes.OpCodeNoop{})
	for fileScanner.Scan() {
		line := fileScanner.Text()
		err := e.ExecuteLine(line)
		if err != nil {
			panic(err)
		}
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	var score int64
	for _, cycleOneBased := range []int64{20, 60, 100, 140, 180, 220} {
		cycleStrength := e.SignalStrength(cycleOneBased)
		fmt.Printf("cycle %d, cycleStrength: %d\n", cycleOneBased, cycleStrength)
		score += cycleStrength
	}
	fmt.Printf("score: %d\n", score)
}
