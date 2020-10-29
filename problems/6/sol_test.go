package solution

import (
	"testing"
)

func TestSol(t *testing.T) {

	testcases := []struct {
		s       string
		numRows int
		ans     string
	}{
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			ans:     "PINALSIGYAHRPI",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			ans:     "PAHNAPLSIIGYIR",
		},
		{
			s:       "A",
			numRows: 1,
			ans:     "A",
		},
		{
			s:       "A",
			numRows: 2,
			ans:     "A",
		},
	}
	for _, testcase := range testcases {
		r := convert(testcase.s, testcase.numRows)
		if r != testcase.ans {
			t.Fatal("wrong result:", r, "expected:", testcase.ans)
		}
	}

	t.Log("... passed")
}
