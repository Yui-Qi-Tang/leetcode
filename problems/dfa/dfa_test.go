package dfa

import (
	"testing"
)

func TestBasic(t *testing.T) {

	// check input with {a+b}
	m := M{
		states:    states{"q0", "q1", "q2"},
		alphabet:  alphabet{"a", "b"},
		acceptSet: states{"q2"},
		initState: "q0",
	}
	m.init()
	input := []string{"a", "a", "a", "b"} // a+b

	for _, v := range input {
		prevState := m.getCurrentState()
		if err := m.transition(v); err != nil {
			t.Fatal(err)
		}
		t.Logf("transition(%s, '%s') -> {%s}", prevState, v, m.getCurrentState())

	}
	if m.isAccept() {
		// t.Logf("%s is in accept set: %+v", m.getCurrentState(), m.acceptSet)
		t.Log("valid input")
	} else {
		t.Fatal("invalid input")
	}
	t.Log("... Passed")

}
