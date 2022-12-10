package engine

import (
	"AOC2022/cmd/day-10/engine/machine"
	"errors"
	"fmt"
)

type Engine struct {
	m        machine.MachineState
	recorder machine.MachineStateRecorder
	opCodes  []OpCode
}

func (e *Engine) ExecuteLine(line string) error {
	for _, opCode := range e.opCodes {
		matches := opCode.Regex().FindAllStringSubmatch(line, -1)
		err := opCode.LoadInstructionImmediates(matches)
		if errors.Is(err, ErrInvalidLoadValues) {
			// It's not this opCode
			continue
		} else if err != nil {
			return err
		}
		subCycles := opCode.SubCycles()
		for i := 0; i < subCycles; i++ {
			e.recorder.RecordDuring(e.m.Cycles, e.m)

			cycleOneBased := e.m.Cycles + 1
			cycleStrength := e.SignalStrength(cycleOneBased)
			fmt.Printf("cycle(one-based) %d, X: %d, cycleStrength: %d, line: %s\n", cycleOneBased, e.m.X, cycleStrength, line)

			err = opCode.Do(i, &e.m)
			if err != nil {
				return err
			}
			e.recorder.RecordAfter(e.m.Cycles, e.m)
			e.m.Cycle()
		}
		break
	}
	return nil
}

func NewEngine() *Engine {
	return &Engine{
		m:        *machine.NewMachineState(),
		recorder: *machine.NewRecorder(),
	}
}

func (e *Engine) AddOpCode(o OpCode) {
	e.opCodes = append(e.opCodes, o)
}

func (e *Engine) SignalStrength(cycle int64) int64 {
	// cycle must be a one-based cycle number ie "cycle one" -> 0th index
	states := e.recorder.DuringStates()
	// Cycles are stored zero indexed. Furthermore, "SignalStrength" is defined as (one-based) cycle * X _during_ cycle.
	state, ok := states[cycle-1]
	if !ok {
		return 0
	}
	return state.X * cycle
}

func (e *Engine) RenderCRT() {
	var lineNum int64
	for lineNum = 0; lineNum < 6; lineNum++ {
		var i int64
		for i = 0; i < 40; i++ {
			cycle := lineNum*40 + i // zero based
			state := e.recorder.DuringStates()[int64(cycle)]
			if (state.X-1) <= i && i <= (state.X+1) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
