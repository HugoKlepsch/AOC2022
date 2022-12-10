package engine

import (
	"AOC2022/cmd/day-10/engine/machine"
	"errors"
	"regexp"
)

var (
	ErrInvalidSubCycle   = errors.New("invalid subCycle given to OpCode")
	ErrInvalidLoadValues = errors.New("invalid subCycle given to OpCode")
)

type OpCode interface {
	Do(subCycle int, m *machine.MachineState) error
	LoadInstructionImmediates(matches [][]string) error
	Regex() *regexp.Regexp
	SubCycles() int
}
