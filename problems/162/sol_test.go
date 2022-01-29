package sol

import "testing"

func TestBasic(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{9, 1},
			want:  0,
		},
		{
			input: []int{1, 9, 1},
			want:  1,
		},
		{
			input: []int{1, 2},
			want:  1,
		},
		{
			input: []int{1, 1, 1, 3, 1},
			want:  3,
		},
		{
			input: []int{1, 2, 3, 1},
			want:  2,
		},
		{
			input: []int{1, 2, 1, 3, 5, 6, 4}, // 1 or 5 are valid
			want:  5,
		},
	}

	for _, tt := range testcases {
		ans := findPeakElementv2(tt.input)
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}
}
