package solution

import (
	"testing"
)

func TestBasic(t *testing.T) {

	testcases := []struct {
		n    int
		want int
	}{
		{n: 0, want: 1},
		{n: 1, want: 1},
		{n: 2, want: 2},
		{n: 3, want: 3},
		{n: 4, want: 5},
		{n: 5, want: 8},
		{n: 6, want: 13},
		{n: 7, want: 21},
		{n: 8, want: 34},
		{n: 9, want: 55},
		{n: 100, want: 1298777728820984005},
	}

	for _, tt := range testcases {
		answer := climbStairs(tt.n)
		if answer != tt.want {
			t.Fatal("it should be:", tt.want, "but got:", answer)
		}
	}

	t.Log("... Passed")
}
