package main

import (
	"bufio"
	"fmt"
	"os"
)

type Selection int

type GameEndState int

const (
	SelectionRock Selection = iota
	SelectionPaper 
	SelectionScissors
	GameEndStateWin GameEndState = iota
	GameEndStateDraw 
	GameEndStateLoss 
)

type StrategyGuideLine struct {
	Opponent Selection
	Recommendation GameEndState
}

var asciiToSelection map[string]Selection

var asciiToGameEndState map[string]GameEndState

var selectionWinner map[Selection]Selection

var selectionLoser map[Selection]Selection

var selectionPoints map[Selection]int

func main() {
	asciiToSelection = map[string]Selection{
		"A": SelectionRock,
		"B": SelectionPaper,
		"C": SelectionScissors,
	}

	asciiToGameEndState = map[string]GameEndState{
		"X": GameEndStateLoss,
		"Y": GameEndStateDraw,
		"Z": GameEndStateWin,
	}

	selectionWinner = map[Selection]Selection{
		SelectionRock: SelectionPaper,
		SelectionPaper: SelectionScissors,
		SelectionScissors: SelectionRock,
	}

	selectionLoser = map[Selection]Selection{
		SelectionRock: SelectionScissors,
		SelectionPaper: SelectionRock,
		SelectionScissors: SelectionPaper,
	}

	selectionPoints = map[Selection]int{
		SelectionRock: 1,
		SelectionPaper: 2,
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
			Opponent: asciiToSelection[opponent],
			Recommendation: asciiToGameEndState[recommendation],
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

	var selectionForRecommendation Selection
	if strategy.Recommendation == GameEndStateLoss {
		// We were told to lose
		selectionForRecommendation = selectionLoser[strategy.Opponent]
	} else if strategy.Recommendation == GameEndStateDraw {
		// We were told to draw
		selectionForRecommendation = strategy.Opponent
	} else {
		// We were told to win
		selectionForRecommendation = selectionWinner[strategy.Opponent]
	}

	if strategy.Opponent == selectionForRecommendation {
		// Draw
		score += 3
	} else if selectionWinner[strategy.Opponent] == selectionForRecommendation {
		// We won
		score += 6
	} else {
		// We lost
		score += 0 // for completeness
	}
	score += selectionPoints[selectionForRecommendation]
	
	return score
}
