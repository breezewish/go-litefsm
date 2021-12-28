package litefsm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransitions(t *testing.T) {
	assert := require.New(t)

	transitions := NewTransitions()
	transitions.AddTransitFrom("CmdDispatchBegin").
		Into("CmdQueryBegin").
		ThenInto("CmdQueryFinish").
		ThenInto("CmdDispatchFinish").
		ThenInto("CmdDispatchBegin")

	assert.EqualError(transitions.CanTransit("Foo", "Bar"), `Foo[x] -> Bar: invalid transit`)
	assert.ErrorIs(transitions.CanTransit("Foo", "Bar"), ErrInvalidTransit)

	assert.EqualError(transitions.CanTransit("CmdDispatchBegin", "CmdDispatchBegin"), `CmdDispatchBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("CmdDispatchBegin", "CmdQueryBegin"))
	assert.EqualError(transitions.CanTransit("CmdDispatchBegin", "CmdQueryFinish"), `CmdDispatchBegin -> CmdQueryFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdDispatchBegin", "CmdDispatchFinish"), `CmdDispatchBegin -> CmdDispatchFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdDispatchBegin", "CmdDispatchBegin"), `CmdDispatchBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdDispatchBegin", "StmtHandleBegin"), `CmdDispatchBegin -> StmtHandleBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdDispatchBegin", "StmtHandleFinish"), `CmdDispatchBegin -> StmtHandleFinish[x]: invalid transit`)

	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdDispatchBegin"), `CmdQueryBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdQueryBegin"), `CmdQueryBegin -> CmdQueryBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("CmdQueryBegin", "CmdQueryFinish"))
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdDispatchFinish"), `CmdQueryBegin -> CmdDispatchFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdDispatchBegin"), `CmdQueryBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "StmtHandleBegin"), `CmdQueryBegin -> StmtHandleBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "StmtHandleFinish"), `CmdQueryBegin -> StmtHandleFinish[x]: invalid transit`)

	assert.EqualError(transitions.CanTransit("CmdQueryFinish", "CmdDispatchBegin"), `CmdQueryFinish -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryFinish", "CmdQueryBegin"), `CmdQueryFinish -> CmdQueryBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryFinish", "CmdQueryFinish"), `CmdQueryFinish -> CmdQueryFinish[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("CmdQueryFinish", "CmdDispatchFinish"))
	assert.EqualError(transitions.CanTransit("CmdQueryFinish", "CmdDispatchBegin"), `CmdQueryFinish -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryFinish", "StmtHandleBegin"), `CmdQueryFinish -> StmtHandleBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryFinish", "StmtHandleFinish"), `CmdQueryFinish -> StmtHandleFinish[x]: invalid transit`)

	transitions.AddTransitFrom("CmdQueryBegin").
		Into("StmtHandleBegin").
		ThenInto("StmtHandleFinish").
		ThenInto("CmdQueryFinish")

	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdQueryBegin"), `CmdQueryBegin -> CmdQueryBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdDispatchBegin"), `CmdQueryBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("CmdQueryBegin", "CmdQueryFinish"))
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdDispatchFinish"), `CmdQueryBegin -> CmdDispatchFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "CmdDispatchBegin"), `CmdQueryBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("CmdQueryBegin", "StmtHandleBegin"))
	assert.EqualError(transitions.CanTransit("CmdQueryBegin", "StmtHandleFinish"), `CmdQueryBegin -> StmtHandleFinish[x]: invalid transit`)

	assert.EqualError(transitions.CanTransit("StmtHandleBegin", "CmdDispatchBegin"), `StmtHandleBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleBegin", "CmdQueryBegin"), `StmtHandleBegin -> CmdQueryBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleBegin", "CmdQueryFinish"), `StmtHandleBegin -> CmdQueryFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleBegin", "CmdDispatchFinish"), `StmtHandleBegin -> CmdDispatchFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleBegin", "CmdDispatchBegin"), `StmtHandleBegin -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleBegin", "StmtHandleBegin"), `StmtHandleBegin -> StmtHandleBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("StmtHandleBegin", "StmtHandleFinish"))

	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdDispatchBegin"), `StmtHandleFinish -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdQueryBegin"), `StmtHandleFinish -> CmdQueryBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("StmtHandleFinish", "CmdQueryFinish"))
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdDispatchFinish"), `StmtHandleFinish -> CmdDispatchFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdDispatchBegin"), `StmtHandleFinish -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "StmtHandleBegin"), `StmtHandleFinish -> StmtHandleBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "StmtHandleFinish"), `StmtHandleFinish -> StmtHandleFinish[x]: invalid transit`)

	transitions.AddTransitFrom("StmtHandleFinish").Into("StmtHandleBegin")

	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdDispatchBegin"), `StmtHandleFinish -> CmdDispatchBegin[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdQueryBegin"), `StmtHandleFinish -> CmdQueryBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("StmtHandleFinish", "CmdQueryFinish"))
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdDispatchFinish"), `StmtHandleFinish -> CmdDispatchFinish[x]: invalid transit`)
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "CmdDispatchBegin"), `StmtHandleFinish -> CmdDispatchBegin[x]: invalid transit`)
	assert.Nil(transitions.CanTransit("StmtHandleFinish", "StmtHandleBegin"))
	assert.EqualError(transitions.CanTransit("StmtHandleFinish", "StmtHandleFinish"), `StmtHandleFinish -> StmtHandleFinish[x]: invalid transit`)

	assert.EqualError(transitions.CanTransit("abc", "CmdDispatchBegin"), `abc[x] -> CmdDispatchBegin: invalid transit`)
	assert.EqualError(transitions.CanTransit("abc", "CmdQueryBegin"), `abc[x] -> CmdQueryBegin: invalid transit`)
	assert.EqualError(transitions.CanTransit("abc", "CmdQueryFinish"), `abc[x] -> CmdQueryFinish: invalid transit`)
	assert.EqualError(transitions.CanTransit("abc", "CmdDispatchFinish"), `abc[x] -> CmdDispatchFinish: invalid transit`)
	assert.EqualError(transitions.CanTransit("abc", "CmdDispatchBegin"), `abc[x] -> CmdDispatchBegin: invalid transit`)
	assert.EqualError(transitions.CanTransit("abc", "StmtHandleBegin"), `abc[x] -> StmtHandleBegin: invalid transit`)
	assert.EqualError(transitions.CanTransit("abc", "StmtHandleFinish"), `abc[x] -> StmtHandleFinish: invalid transit`)
}
