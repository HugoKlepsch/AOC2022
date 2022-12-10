package opcodes

import (
	"AOC2022/cmd/day-10/engine"
	"AOC2022/cmd/day-10/engine/machine"
	"fmt"
	"regexp"
	"strconv"
)

type OpCodeAddX struct {
	regex     *regexp.Regexp
	immediate int64
}

func (o *OpCodeAddX) Do(subCycle int, m *machine.MachineState) error {
	switch subCycle {
	case 0:
		return nil
	case 1:
		m.X += o.immediate
		return nil
	default:
		return fmt.Errorf("%w: %d", engine.ErrInvalidSubCycle, subCycle)
	}
}

func (o *OpCodeAddX) LoadInstructionImmediates(matches [][]string) error {
	if matches == nil || len(matches) != 1 || len(matches[0]) != 2 || matches[0][0] == "" || matches[0][1] == "" {
		return fmt.Errorf("%w: %v", engine.ErrInvalidLoadValues, matches)
	}

	i, err := strconv.ParseInt(matches[0][1], 10, 64)

	if err != nil {
		return err
	}
	o.immediate = i
	return nil
}

func (o *OpCodeAddX) Regex() *regexp.Regexp {
	if o.regex == nil {
		o.regex = regexp.MustCompile(`^addx (-?[0-9]+)$`)
	}
	return o.regex
}

func (o *OpCodeAddX) SubCycles() int {
	return 2
}
