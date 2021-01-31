package sol

import "testing"

func TestBasic(t *testing.T) {

	testcases := []struct {
		input []int
		want  [][]int
	}{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for _, testcase := range testcases {
		result := threeSum(testcase.input)
		if len(result) != len(testcase.want) {
			t.Fatal("expected length is:", len(testcase.want), "but got:", len(result))
		}

		for i, want := range testcase.want {
			res := result[i]

			for j, w := range want {
				if res[j] != w {
					t.Fatal("expected:", testcase.want, "but got:", result)
				}
			}
		}
	}

	t.Log("passed...")
}
