package litefsm

type StateMachine struct {
	transitions  Transitions
	currentState State
}

func NewStateMachine(transitions Transitions, initialState State) *StateMachine {
	return &StateMachine{
		transitions:  transitions,
		currentState: initialState,
	}
}

func (sm *StateMachine) Current() State {
	return sm.currentState
}

func (sm *StateMachine) ResetTo(state State) {
	sm.currentState = state
}

func (sm *StateMachine) CanGoto(newState State) error {
	return sm.transitions.CanTransit(sm.currentState, newState)
}

func (sm *StateMachine) Goto(newState State) error {
	err := sm.CanGoto(newState)
	if err != nil {
		return err
	}
	sm.currentState = newState
	return nil
}
