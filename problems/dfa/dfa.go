package dfa

import (
	"errors"
)

// M DFA model
type M struct {
	states       states
	alphabet     alphabet
	acceptSet    states
	initState    string
	currentState string
}

type alphabet []string

func (a alphabet) exist(input string) bool {
	for _, v := range a {
		if v == input {
			return true
		}
	}
	return false
}

type states []string

func (s states) exist(input string) bool {
	for _, v := range s {
		if v == input {
			return true
		}
	}
	return false
}

func (m *M) init() {
	m.currentState = m.initState
}

func (m *M) transition(input string) error {
	if m.alphabet.exist(input) == false {
		return errors.New(input + " does not exist in alphbat")
	}
	if m.currentState == "q0" && input == "a" {
		m.currentState = "q1"
	} else if m.currentState == "q1" && input == "a" {
		m.currentState = "q1"
	} else if m.currentState == "q1" && input == "b" {
		m.currentState = "q2"
	} else {
		return errors.New("unknown transition")
	}
	return nil

}

func (m M) getCurrentState() string {
	return m.currentState
}

func (m M) isAccept() bool {

	return m.acceptSet.exist(m.getCurrentState())
}
