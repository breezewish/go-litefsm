package litefsm

import (
	"errors"
	"fmt"
)

type State string

type Transitions map[State]map[State]struct{} // FromState -> ToState

var ErrInvalidTransit = errors.New("invalid transit")

func NewTransitions() Transitions {
	return make(map[State]map[State]struct{})
}

func (t Transitions) AddTransit(fromState, intoState State) {
	target, ok := t[fromState]
	if !ok {
		target = make(map[State]struct{})
		t[fromState] = target
	}
	target[intoState] = struct{}{}
}

type TransitFrom struct {
	t         Transitions
	fromState State
}

type TransitInto struct {
	t         Transitions
	fromState State
}

func (t Transitions) AddTransitFrom(state State) TransitFrom {
	return TransitFrom{
		t:         t,
		fromState: state,
	}
}

func (f TransitFrom) Into(state State) TransitInto {
	f.t.AddTransit(f.fromState, state)
	return TransitInto{
		t:         f.t,
		fromState: state,
	}
}

func (i TransitInto) ThenInto(state State) TransitInto {
	i.t.AddTransit(i.fromState, state)
	return TransitInto{
		t:         i.t,
		fromState: state,
	}
}

func (t Transitions) CanTransit(fromState, intoState State) error {
	target, ok := t[fromState]
	if !ok {
		return fmt.Errorf("%s[x] -> %s: %w", fromState, intoState, ErrInvalidTransit)
	}
	_, ok = target[intoState]
	if !ok {
		return fmt.Errorf("%s -> %s[x]: %w", fromState, intoState, ErrInvalidTransit)
	}

	return nil
}
