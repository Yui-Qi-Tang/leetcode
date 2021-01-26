package sol

import (
	"testing"
)

/*
Constraints:
  n == height.length
  2 <= n <= 3 * 104
  0 <= height[i] <= 3 * 104
*/

func TestBasic(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{input: []int{1, 1}, want: 1},
		{input: []int{4, 3, 2, 1, 4}, want: 16},
		{input: []int{1, 2, 1}, want: 2},
	}

	for _, testcase := range testcases {
		result := maxArea(testcase.input)
		if result != testcase.want {
			t.Fatal("expected:", testcase.want, "but got:", result)
		}
	}

	t.Log("passed")
}
