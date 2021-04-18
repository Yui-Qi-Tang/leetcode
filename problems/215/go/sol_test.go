package sol

import "testing"

func TestMaxHeapify(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 2, 3, 4, 100},
			want:  100,
		},
		{
			input: []int{1, 2, 100, 4, 5},
			want:  100,
		},
		{
			input: []int{100, 2, 3, 4, 5},
			want:  100,
		},
	}

	for _, tt := range testcases {
		// build max heapify
		for i := len(tt.input) >> 1; i >= 0; i-- {
			maxHeapify(tt.input, i)
		}
		if tt.input[0] != tt.want {
			t.Fatal("it should be:", tt.want, "but got:", tt.input[0])
		}
	}

	t.Log("... Passed")

}

func TestFindKthLargest(t *testing.T) {
	testcases := []struct {
		input []int
		k     int
		want  int
	}{
		{
			input: []int{3, 2, 1, 5, 6, 4},
			k:     2,
			want:  5,
		},
		{
			input: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:     4,
			want:  4,
		},
		{
			input: []int{2, -1, 3},
			k:     3,
			want:  -1,
		},
	}

	for _, tt := range testcases {
		res := findKthLargest(tt.input, tt.k)

		if res != tt.want {
			t.Fatal("it should be", tt.want, "but got", res)
		}
	}
}
