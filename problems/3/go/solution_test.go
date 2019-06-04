package solution

import "testing"

type testCase struct {
	answer int
	input string
}

func TestSolution(t *testing.T) {
	cases := []testCase{
		testCase{
			answer: 3,
			input : "abcabcbb",
		},
		testCase{
			answer: 3,
			input : "testtest",
		},
		testCase{
			answer: 3,
			input : "pwwkew",
		},
		testCase{
			answer: 1,
			input : "bbbbbbb",
		},
		testCase{
			answer: 0,
			input : "",
		},
		testCase{
			answer: 1,
			input : " ",
		},
		testCase{
			answer: 2,
			input : "ab",
		},
		testCase{
			answer: 3,
			input : "ab ab",
		},
		testCase{
			answer: 4,
			input : "abcd",
		},
		testCase{
			answer: 2,
			input : "aab",
		},
		testCase{
			answer: 6,
			input : "asjrgapa",
		},
	}
	for _, c := range cases {
		answer := lengthOfLongestSubstring(c.input)
		t.Logf("Start testing %s, correct answer is: %d; your answer is: %d", c.input, c.answer, answer)
		if c.answer != answer {
			t.Fatal("Got wrong answer!")
		}else {
			t.Log(" ...PASS")
		}
	}
}