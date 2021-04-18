package quicksort

import (
	"testing"
)

func TestPartition(t *testing.T) {
	a := []int{1, 2, 5, 4, 3}
	t.Log(partition(a, 0, len(a)-1))

}

func TestSort(t *testing.T) {
	a := []int{2, 8, 7, 1, 3, 5, 6, 4}
	sort(a, 0, len(a)-1)
	t.Log(a)
}

func Benchmark1(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		sort(a, 0, len(a)-1)
	}
}

func Benchmark2(b *testing.B) {
	a := []int{2, 8, 7, 1, 3, 5, 6, 4}
	for i := 0; i < b.N; i++ {
		sort(a, 0, len(a)-1)
	}
}

func Benchmark3(b *testing.B) {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < b.N; i++ {
		sort(a, 0, len(a)-1)
	}
}
