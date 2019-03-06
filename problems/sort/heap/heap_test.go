package heap

import (
	"math/rand"
	"testing"
	"time"
)

// for create a big size int array
const size int = 1000000

func createTestData(arraySize int) []int {
	var list = make([]int, arraySize)
	var prng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := range list {
		list[i] = prng.Int()
	}
	return list
}

func TestCorrectness(t *testing.T) {
	A := createTestData(size) // A[0] is not used!
	buildMaxHeap(A)
	heapSort(A, false)
	for i := 1; i < len(A); i++ {
		next := i + 1
		if next >= len(A) {
			break
		}
		if A[i] > A[next] {
			t.Fatal("Failed!")
		}
	}
}

func BenchmarkHeap(b *testing.B) {

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		list := createTestData(size)
		b.StartTimer()
		heapSort(list, false)
	}
}
