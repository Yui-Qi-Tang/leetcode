package sol

import (
	"testing"
)

func TestPow(t *testing.T) {

	testcases := []struct {
		base, pos, want int
	}{
		{base: 10, pos: 1, want: 10},
		{base: 10, pos: 2, want: 100},
		{base: -1, pos: 2, want: 1},
		{base: -1, pos: 3, want: -1},
	}

	for _, testcase := range testcases {
		result := pow(testcase.base, testcase.pos)
		if result != testcase.want {
			t.Fatal("expected:", testcase.want, "but got:", result)
		}
	}

	t.Log("passed")
	// intToRoman(1551)
}

func TestBasic(t *testing.T) {

	testcases := []struct {
		input int
		want  string
	}{
		{input: 3, want: "III"},
		{input: 4, want: "IV"},
		{input: 9, want: "IX"},
		{input: 58, want: "LVIII"},
		{input: 1994, want: "MCMXCIV"},
		{input: 1551, want: "MDLI"},
	}

	for _, testcase := range testcases {
		result := intToRoman(testcase.input)
		if result != testcase.want {
			t.Fatal("expected:", testcase.want, "but got:", result)
		}
	}
}
