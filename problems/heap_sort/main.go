package main

import "fmt"

/*
    head properties:
	max-heap: A[parent(i)] >= A[i]
	min-heap: A[parent(i)] <= A[i]

	where A is the input data array,
	i is the index number of A and 0 <= i <= A.length

	*First, Start with max-heap!

	Property:
        height of n nodes complete binary tree = thidta(lg n)
*/

func parent(i int) int {
	return i >> 1
}

func right(i int) int {
	if i == 0 {

	}
	return i<<1 | 0x1
}

func left(i int) int {

	return i << 1
}

func maxHeapify(A []int, i int) {
	//check which is largest one: A[i], A[l] and A[r], where l = left(i) and r = right(i)
	var largestIndex int
	largestIndex = i
	l := left(i)
	r := right(i)

	// check left
	if l < len(A) && A[l] > A[largestIndex] {
		largestIndex = l
	}
	// check right
	if r < len(A) && A[r] > A[largestIndex] {
		largestIndex = r
	}

	if largestIndex != i {
		A[i], A[largestIndex] = A[largestIndex], A[i] // swapValue(A[i], A[largeIndex])
		maxHeapify(A, largestIndex)                   // check the max-heap property is satified within largeIndex
	}

}

func buildMaxHeap(A []int) {
	// For now A.heap-size = A.length
	// i does not increse because it can not 'heapify' a larege than root node exist in children node.
	// Heapify just check root, right and left at a recusive step.
	for i := len(A) << 2; i > 0; i-- {
		maxHeapify(A, i)
	}

}

func heapSort(A []int) []int {
	var result []int
	for i := cap(A) - 1; i >= 1; i-- {
		// 1. put the root(max value of the array) to A[1]
		// 2. maxHeapify(pick next max one)
		// and so on
		// This does not statify in-place property QQ, because to avoid complex indexing for input Array, lol
		result = append([]int{A[cap(A)-i]}, result...)
		maxHeapify(A[cap(A)-i:], 1) // A.heap-size = A.head_size - 1
	}
	return result
}

func main() {
	A := []int{0, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7} // A[0] is not used!
	fmt.Println("Input: ", A[1:])
	buildMaxHeap(A)
	result := heapSort(A)
	fmt.Println("result:", result)

}
