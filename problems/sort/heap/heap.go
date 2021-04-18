package heap

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

func heapSort(A []int, showProcess bool) {
	if showProcess {
		fmt.Println("displaying sort proccess... ")
	}

	// cap(A) - 1: skip A[0]; i: heap-size
	for i := cap(A) - 1; i >= 1; i-- {
		if showProcess {
			fmt.Println("Input:", A[1:], "heap size:", i)
		}
		A[1], A[i] = A[i], A[1] // let min value on heap tree root <> max-heap property
		maxHeapify(A[:i], 1)    // A[:i]: decrease head-size via i--;  'float down' the min value XD and let the max value left the heap tree
	}
}
