package sol

import "testing"

func TestBasic(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{},
			want:  -1,
		},
		{
			input: []int{1, 2, 3},
			want:  1,
		},
		{
			input: []int{2, 1},
			want:  1,
		},
		{
			input: []int{100},
			want:  100,
		},
		{
			input: []int{3, 4, 5, 1, 2},
			want:  1,
		},
		{
			input: []int{4, 5, 6, 7, 0, 1, 2},
			want:  0,
		},
		{
			input: []int{11, 13, 15, 17},
			want:  11,
		},
		{
			input: []int{5, 1, 2, 3, 4},
			want:  1,
		},
	}

	for i, tt := range testcases {
		t.Log("case", i)
		ans := findMin(tt.input)
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}
}
