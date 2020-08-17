package solution

import (
	"strconv"
	"testing"
)

func TestSanbox(t *testing.T) {
	pos := 123

	neg := -123

	posStr := "012"

	t.Log("pos int to string:", strconv.Itoa(pos))

	t.Log("neg int to string:", strconv.Itoa(neg))

	v, err := strconv.Atoi(posStr)
	if err != nil {
		t.Log(err)
	}

	t.Log("posStr string to int:", v)

}

func TestStringReverse(t *testing.T) {

	posStr := "120"
	posStrR := reverseString(posStr)

	v, err := strconv.Atoi(posStr)
	if err != nil {
		t.Log(err)
	}

	t.Log("posStr string to int:", v)

	v2, err := strconv.Atoi(posStrR)
	if err != nil {
		t.Log(err)
	}

	t.Log("posStrR string to int:", v2)

	t.Log("max:", 1<<31-1, "min:", -1<<31)

}

func TestSolution(t *testing.T) {
	testCases := []struct {
		input  int
		answer int
	}{
		{123, 321}, {-123, -321}, {120, 21},
		{1534236469, 0}, {-2147483648, 0}, {1463847412, 2147483641},
	}

	for _, testcase := range testCases {
		v := reverse(testcase.input)

		if v != testcase.answer {
			t.Fatal(v, "!=", testcase.answer, "(answer)")
		}

	}

	t.Log("...Passed")
}
