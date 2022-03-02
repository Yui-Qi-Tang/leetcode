package sol

import "testing"

func TestBasic(t *testing.T) {

	testcases := []struct {
		input  []int
		target int
		want   int
	}{
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			input:  []int{1},
			target: 0,
			want:   -1,
		},
		{
			input:  []int{1},
			target: 1,
			want:   0,
		},
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			want:   0,
		},
		{
			input:  []int{5, 6, 7, 0, 1, 2, 4},
			target: 4,
			want:   6,
		},
		{
			input:  []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			input:  []int{1, 3, 9},
			target: 3,
			want:   1,
		},
		{
			input:  []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			input:  []int{1, 3, 5},
			target: 5,
			want:   2,
		},
		{
			input:  []int{1, 3, 4, 5},
			target: 5,
			want:   3,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
		{
			input:  []int{5, 1, 2, 3, 4},
			target: 1,
			want:   1,
		},
	}

	for i, tt := range testcases {
		ans := search(tt.input, tt.target)
		if ans != tt.want {
			t.Fatalf("cases[%d]: it should be %d, but got %d", i, tt.want, ans)
		}
	}

}
