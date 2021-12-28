package litefsm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFSM(t *testing.T) {
	assert := require.New(t)

	transitions := NewTransitions()
	transitions.AddTransitFrom("CmdDispatchBegin").
		Into("CmdQueryBegin").
		ThenInto("CmdQueryFinish").
		ThenInto("CmdDispatchFinish").
		ThenInto("CmdDispatchBegin")
	transitions.AddTransitFrom("CmdQueryBegin").
		Into("StmtHandleBegin").
		ThenInto("StmtHandleFinish").
		ThenInto("CmdQueryFinish")
	transitions.AddTransitFrom("StmtHandleFinish").
		Into("StmtHandleBegin")

	fsm := NewStateMachine(transitions, "foo")
	assert.EqualValues("foo", fsm.Current())
	assert.ErrorIs(fsm.Goto("CmdDispatchBegin"), ErrInvalidTransit)
	assert.EqualValues("foo", fsm.Current())

	fsm = NewStateMachine(transitions, "CmdQueryBegin")
	assert.EqualValues("CmdQueryBegin", fsm.Current())
	assert.Nil(fsm.CanGoto("StmtHandleBegin"))
	assert.EqualValues("CmdQueryBegin", fsm.Current())
	assert.Nil(fsm.CanGoto("CmdQueryFinish"))
	assert.EqualValues("CmdQueryBegin", fsm.Current())
	assert.ErrorIs(fsm.CanGoto("StmtHandleFinish"), ErrInvalidTransit)
	assert.EqualValues("CmdQueryBegin", fsm.Current())

	assert.Nil(fsm.Goto("StmtHandleBegin"))
	assert.Nil(fsm.Goto("StmtHandleFinish"))
	assert.Nil(fsm.Goto("StmtHandleBegin"))
	assert.Nil(fsm.Goto("StmtHandleFinish"))
	assert.Nil(fsm.Goto("CmdQueryFinish"))
}
