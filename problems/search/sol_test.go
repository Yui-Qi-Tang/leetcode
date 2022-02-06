package sol

import "testing"

func TestSearchLinear(t *testing.T) {

	testcases := []struct {
		in   []uint
		n    uint
		want int
	}{
		{
			in:   []uint{1, 2, 3},
			n:    1,
			want: 1,
		},
		{
			in:   []uint{1, 2, 3},
			n:    5,
			want: -1,
		},
	}

	for _, tt := range testcases {
		ans := searchLinear(tt.in, uint(tt.n))
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}

}

func TestSearchDNC(t *testing.T) {

	testcases := []struct {
		in   []uint
		n    uint
		want int
	}{
		{
			in:   []uint{1, 2, 3},
			n:    1,
			want: 1,
		},
		{
			in:   []uint{1, 2, 3},
			n:    5,
			want: -1,
		},
		{
			in:   []uint{3, 2, 1},
			n:    1,
			want: 1,
		},
	}

	for _, tt := range testcases {
		ans := searchDNC(tt.in, uint(tt.n))
		if ans != tt.want {
			t.Fatalf("it should be %d, but got %d", tt.want, ans)
		}
	}

}

func BenchmarkSearchLinear(b *testing.B) {
	in := []uint{
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
	}
	var n uint = 9

	for i := 0; i < b.N; i++ {
		searchLinear(in, n)
	}
}

func BenchmarkSearchDNC(b *testing.B) {
	in := []uint{
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
	}
	var n uint = 9

	for i := 0; i < b.N; i++ {
		searchDNC(in, n)
	}
}
