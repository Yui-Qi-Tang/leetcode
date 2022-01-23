package sol

import "testing"

func TestBasic(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{7, 1, 5, 3, 6, 4},
			want:  5,
		},
		{
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
		{
			input: []int{},
			want:  0,
		},
		{
			input: []int{2, 1, 4},
			want:  3,
		},
		{
			input: []int{2, 4, 1},
			want:  2,
		},
		{
			input: []int{3, 2, 6, 5, 0, 3},
			want:  4,
		},
	}

	for i, tt := range testcases {
		ans := maxProfit4(tt.input)
		if ans != tt.want {
			t.Fatalf("case %d it should be %d, but got: %d", i, tt.want, ans)
		}
	}

}
