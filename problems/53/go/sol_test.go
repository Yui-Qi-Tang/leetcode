package sol

import (
	"testing"
)

func TestBasic(t *testing.T) {

	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want:  6,
		},
		{
			input: []int{1},
			want:  1,
		},
		{
			input: []int{-1, -2, -3},
			want:  -1,
		},
		{
			input: []int{-1, -2, -3, 1, -1, 9},
			want:  9,
		},
	}

	for _, tt := range testcases {
		answer := maxSubArray(tt.input)

		if answer != tt.want {
			t.Fatal("it should be:", tt.want, "but got:", answer)
		}
	}

	t.Log("... Passed")

}
