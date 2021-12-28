package litefsm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDebugGraphviz(t *testing.T) {
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

	v := transitions.DebugGraphviz()
	require.Equal(t, strings.TrimSpace(`
digraph fsm {
    "CmdDispatchBegin" -> "CmdQueryBegin";

    "CmdDispatchFinish" -> "CmdDispatchBegin";

    "CmdQueryBegin" -> "CmdQueryFinish";
    "CmdQueryBegin" -> "StmtHandleBegin";

    "CmdQueryFinish" -> "CmdDispatchFinish";

    "StmtHandleBegin" -> "StmtHandleFinish";

    "StmtHandleFinish" -> "CmdQueryFinish";
    "StmtHandleFinish" -> "StmtHandleBegin";

}
`), strings.TrimSpace(v))
}
