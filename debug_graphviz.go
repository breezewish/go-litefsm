package litefsm

import (
	"bytes"
	"fmt"
	"sort"
)

func (t Transitions) DebugGraphviz() string {
	var buf bytes.Buffer
	writeHeader(&buf)
	writeTransitions(&buf, t)
	writeFooter(&buf)
	return buf.String()
}

func writeHeader(buf *bytes.Buffer) {
	buf.WriteString(`digraph fsm {`)
	buf.WriteString("\n")
}

func writeTransitions(buf *bytes.Buffer, t Transitions) {
	var beginStates []string
	for state := range t {
		beginStates = append(beginStates, string(state))
	}
	sort.Strings(beginStates)
	for _, beginState := range beginStates {
		var endStates []string
		for state := range t[State(beginState)] {
			endStates = append(endStates, string(state))
		}
		sort.Strings(endStates)
		for _, endState := range endStates {
			buf.WriteString(fmt.Sprintf(`    "%s" -> "%s";`, beginState, endState))
			buf.WriteString("\n")
		}
		buf.WriteString("\n")
	}
}

func writeFooter(buf *bytes.Buffer) {
	buf.WriteString(`}`)
	buf.WriteString("\n")
}
