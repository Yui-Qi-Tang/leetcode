package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		input  int
		answer int
	}{
		{123, 321}, {-123, -321}, {120, 21}, // test cases are from 'run code'
		{1534236469, 0}, {-2147483648, 0}, {1463847412, 2147483641}, // test cases are from 'submit'
	}

	for _, testcase := range testCases {
		v := reverse(testcase.input)

		if v != testcase.answer {
			t.Fatal(v, "!=", testcase.answer, "(answer)")
		}

	}

	t.Log("...Passed")
}
