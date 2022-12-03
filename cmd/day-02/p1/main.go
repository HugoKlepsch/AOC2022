package main

import (
	"bufio"
	"fmt"
	"os"
)

type Selection int

const (
	SelectionRock Selection = iota
	SelectionPaper
	SelectionScissors
)

type StrategyGuideLine struct {
	Opponent       Selection
	Recommendation Selection
}

var asciiToSelection map[string]Selection

var selectionWinner map[Selection]Selection

var selectionPoints map[Selection]int

func main() {
	asciiToSelection = map[string]Selection{
		"A": SelectionRock,
		"B": SelectionPaper,
		"C": SelectionScissors,
		"X": SelectionRock,
		"Y": SelectionPaper,
		"Z": SelectionScissors,
	}

	selectionWinner = map[Selection]Selection{
		SelectionRock:     SelectionPaper,
		SelectionPaper:    SelectionScissors,
		SelectionScissors: SelectionRock,
	}

	selectionPoints = map[Selection]int{
		SelectionRock:     1,
		SelectionPaper:    2,
		SelectionScissors: 3,
	}

	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	strategyGuide := make([]StrategyGuideLine, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		opponent := ""
		recommendation := ""

		_, err := fmt.Sscanf(line, "%s %s", &opponent, &recommendation)
		if err != nil {
			panic(err)
		}

		strategyGuideLine := StrategyGuideLine{
			Opponent:       asciiToSelection[opponent],
			Recommendation: asciiToSelection[recommendation],
		}
		strategyGuide = append(strategyGuide, strategyGuideLine)
	}
	if err = fileScanner.Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	score := 0
	for _, strategyGuideLine := range strategyGuide {
		score = score + calcPointsForRound(strategyGuideLine)
	}

	fmt.Println(score)
}

func calcPointsForRound(strategy StrategyGuideLine) int {
	score := 0

	if strategy.Opponent == strategy.Recommendation {
		// Draw
		score += 3
	} else if selectionWinner[strategy.Opponent] == strategy.Recommendation {
		// We won
		score += 6
	} else {
		// We lost
		score += 0 // for completeness
	}
	score += selectionPoints[strategy.Recommendation]

	return score
}
