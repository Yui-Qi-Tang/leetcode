package sol

import "testing"

func TestFindPeakElement(t *testing.T) {
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
		ans := findPeakElement(tt.input)
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}
}

func TestSol1(t *testing.T) {
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
		ans := sol1(tt.input)
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}
}

func TestSol2(t *testing.T) {
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
		ans := sol2(tt.input)
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}
}

func TestSol3(t *testing.T) {
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
		ans := sol3(tt.input)
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}
}

func BenchmarkSol1(b *testing.B) {
	var input []int = []int{1, 2, 1, 3, 5, 6, 4}

	for i := 0; i < b.N; i++ {
		sol1(input)
	}
}

func BenchmarkSol2(b *testing.B) {
	var input []int = []int{1, 2, 1, 3, 5, 6, 4}

	for i := 0; i < b.N; i++ {
		sol2(input)
	}
}

func BenchmarkSol3(b *testing.B) {
	var input []int = []int{1, 2, 1, 3, 5, 6, 4}

	for i := 0; i < b.N; i++ {
		sol3(input)
	}
}
