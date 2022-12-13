package machine

import "fmt"

func NewRecorder() *MachineStateRecorder {
	return &MachineStateRecorder{
		duringStates: make(map[int64]MachineState),
		afterStates:  make(map[int64]MachineState),
	}
}

type MachineStateRecorder struct {
	duringStates map[int64]MachineState
	afterStates  map[int64]MachineState
}

func (m *MachineStateRecorder) RecordDuring(cycle int64, state MachineState) {
	if _, ok := m.duringStates[cycle]; ok {
		err := fmt.Errorf("cycle already recorded: %d", cycle)
		panic(err)
	}

	m.duringStates[cycle] = state
}

func (m *MachineStateRecorder) RecordAfter(cycle int64, state MachineState) {
	if _, ok := m.afterStates[cycle]; ok {
		err := fmt.Errorf("cycle already recorded: %d", cycle)
		panic(err)
	}

	m.afterStates[cycle] = state
}

func (m *MachineStateRecorder) DuringStates() map[int64]MachineState {
	return m.duringStates
}

func (m *MachineStateRecorder) AfterStates() map[int64]MachineState {
	return m.afterStates
}
