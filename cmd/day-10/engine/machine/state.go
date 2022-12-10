package machine

func NewMachineState() *MachineState {
	return &MachineState{
		X:      1,
		Cycles: 0,
	}	
}

type MachineState struct {
	X      int64 // TODO add other registers
	Cycles int64
}

func (m *MachineState) Cycle() {
	m.Cycles++
}
