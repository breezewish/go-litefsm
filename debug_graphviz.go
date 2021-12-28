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

func getSortedTransitBeginStates(t Transitions) []string {
	var beginStates []string
	for state := range t {
		beginStates = append(beginStates, string(state))
	}
	sort.Strings(beginStates)
	return beginStates
}

func writeTransitions(buf *bytes.Buffer, t Transitions) {
	beginStates := getSortedTransitBeginStates(t)
	for _, beginState := range beginStates {
		for endState := range t[State(beginState)] {
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
