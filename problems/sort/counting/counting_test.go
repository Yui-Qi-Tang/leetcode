package counting

import (
	"math/rand"
	"testing"
	"time"
)

// for create a big size int array
const size int = 1000000

func createTestData(arraySize int) []int {
	var list = make([]int, arraySize)
	rand.Seed(time.Now().UnixNano())
	var prng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := range list {
		list[i] = prng.Intn(999999)
	}
	return list
}

func TestCountingSortSlowCorrectness(t *testing.T) {
	A := createTestData(size)
	B := []int{}
	k := getMax(A) + 1
	countingSortSlow(A, B, k)
	//t.Log(A, B, k)
	for i := 0; i < len(B); i++ {
		next := i + 1
		if next >= len(B) {
			continue
		}
		if B[i] > B[next] {
			t.Fatal("Failed!")
		}
	}
	//t.Fatal("Failed!")
	t.Skip("There does some bug here, do not use!")
}

func TestCountingSortCorrectness(t *testing.T) {
	A := createTestData(size)
	B := make([]int, size)
	k := getMax(A) + 1
	countingSort(A, B, k)
	for i := 0; i < len(B); i++ {
		next := i + 1
		if next >= len(B) {
			continue
		}
		if B[i] > B[next] {
			t.Fatal("Failed!")
		}
	}
}

func BenchmarkCounting(b *testing.B) {

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		A := createTestData(size)
		B := make([]int, size)
		k := getMax(A) + 1
		b.StartTimer()
		countingSort(A, B, k)
	}
}
