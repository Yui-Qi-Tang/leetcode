package sol

import (
	"reflect"
	"testing"
)

func TestSol(t *testing.T) {

	testcases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 4},
			want:  []int{24, 12, 8, 6},
		},
		{
			input: []int{-1, 1, 0, -3, 3},
			want:  []int{0, 0, 9, 0, 0},
		},
		{
			input: []int{1, 1},
			want:  []int{1, 1},
		},
		{
			input: []int{0, 0},
			want:  []int{0, 0},
		},
	}

	for _, tt := range testcases {
		ans := productExceptSelf2(tt.input)
		if !reflect.DeepEqual(tt.want, ans) {
			t.Fatalf("it should be: %v, but got: %v", tt.want, ans)
		}
	}

}
