package heap

import (
	"math/rand"
	gosort "sort"
	"testing"
	"time"
)

func randomInts(nums int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(nums)
}

var testcase []int = randomInts(1000)

func sortedSmallestToLargest(array []int) bool {
	for i := 0; i < len(array)-2; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}

	return true
}

func sortedLargestToSmallest(array []int) bool {
	for i := 0; i < len(array)-2; i++ {
		if array[i] < array[i+1] {
			return false
		}
	}

	return true
}

func TestHeapSortWithMaxHeapify(t *testing.T) {
	a := randomInts(1000)
	sort(a)
	if !sortedSmallestToLargest(a) {
		t.Fatal("failed to sort")
	}
	t.Log("passed")
}

func TestMinsort(t *testing.T) {
	a := randomInts(1000)
	minsort(a)
	if !sortedLargestToSmallest(a) {
		t.Fatal("failed to sort")
	}
	t.Log("passed")
}

func BenchmarkSortWithMaxHeapify(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sort(testcase)
	}
}

func BenchmarkSortWithMinHeapify(b *testing.B) {

	for i := 0; i < b.N; i++ {
		minsort(testcase)
	}
}

func BenchmarkGoStdSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		gosort.Ints(testcase)
	}
}
