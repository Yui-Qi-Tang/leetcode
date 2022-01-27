package sol

import "testing"

func TestBasic(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{2, 3, -2, 4},
			want:  6,
		},
		{
			input: []int{2, 3, -2, 4, -5},
			want:  240,
		},
		{
			input: []int{-2, 0, -1},
			want:  0,
		},
		{
			input: []int{2, 3, 2, 4},
			want:  48,
		},
		{
			input: []int{3, 2, 2, 4},
			want:  48,
		},
		{
			input: []int{-2, 3, 2, 4}, // valid 3,2,4 ?
			want:  24,
		},
		{
			input: []int{0, 2},
			want:  2,
		},
	}

	for i, tt := range testcases {
		t.Log("case", i)
		ans := maxProduct(tt.input)
		if ans != tt.want {
			t.Fatalf("testcases[%d] should be %d, but got %d", i, tt.want, ans)
		}
	}
}
