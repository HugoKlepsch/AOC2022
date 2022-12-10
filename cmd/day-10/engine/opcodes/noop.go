package opcodes

import (
	"AOC2022/cmd/day-10/engine"
	"AOC2022/cmd/day-10/engine/machine"
	"fmt"
	"regexp"
)

type OpCodeNoop struct {
	regex *regexp.Regexp
}

func (o *OpCodeNoop) Do(subCycle int, m *machine.MachineState) error {
	switch subCycle {
	case 0:
		return nil
	default:
		return fmt.Errorf("%w: %d", engine.ErrInvalidSubCycle, subCycle)
	}
}

func (o *OpCodeNoop) LoadInstructionImmediates(matches [][]string) error {
	if matches == nil || len(matches) != 1 || len(matches[0]) != 1 || matches[0][0] != "noop" {
		return fmt.Errorf("%w: %v", engine.ErrInvalidLoadValues, matches)
	}
	return nil
}

func (o *OpCodeNoop) Regex() *regexp.Regexp {
	if o.regex == nil {
		o.regex = regexp.MustCompile(`^noop$`)
	}
	return o.regex
}

func (o *OpCodeNoop) SubCycles() int {
	return 1
}
