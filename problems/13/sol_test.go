package sol

import "testing"

func TestBaisc(t *testing.T) {

	testcases := []struct {
		input string
		want  int
	}{
		{input: "III", want: 3},
		// {input: "IV", want: 4},
		// {input: "IX", want: 9},
		// {input: "LVIII", want: 58},
		// {input: "MCMXCIV", want: 1994},
		// {input: "MDLI", want: 1551},
	}

	for _, testcase := range testcases {
		result := romanToInt(testcase.input)

		if result != testcase.want {
			t.Fatal("expected:", testcase.want, "bot got:", result)
		}
	}

	t.Log("passed")

}
