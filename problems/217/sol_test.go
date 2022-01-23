package sol

import "testing"

/*
Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func TestSol(t *testing.T) {
	testcases := []struct {
		input []int
		want  bool
	}{
		{
			input: []int{1, 2, 3, 1},
			want:  true,
		},
		{
			input: []int{1, 2, 3, 4},
			want:  false,
		},
		{
			input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want:  true,
		},
	}

	for i, tt := range testcases {
		ans := containsDuplicate(tt.input)
		if ans != tt.want {
			t.Fatalf("case: %d => it should be %v, but got %v", i, tt.want, ans)
		}
	}
}
