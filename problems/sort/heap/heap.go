package heap

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

// floor(i/2)
func parent(i int) int {
	return i / 2
}

// 2i
func left(i int) int {
	return i << 1
}

// 2i+1
func right(i int) int {
	return i<<1 | 0x1
}

// max-heap: A[parent(i)] >= A[i]
func maxHeapify(array []int, i int) {
	largestIdx := i

	l := left(i)
	r := right(i)

	if l < len(array) && array[l] > array[i] {
		largestIdx = l
	}

	if r < len(array) && array[r] > array[largestIdx] {
		largestIdx = r
	}

	if i != largestIdx {
		array[i], array[largestIdx] = array[largestIdx], array[i]
		maxHeapify(array, largestIdx)
	}
}

// min-heap: A[parent(i)] <= A[i]
func minHeapify(array []int, i int) {
	smallest := i

	l := left(i)
	r := right(i)

	if l < len(array) && array[l] < array[i] {
		smallest = l
	}

	if r < len(array) && array[r] < array[smallest] {
		smallest = r
	}

	if i != smallest {
		array[i], array[smallest] = array[smallest], array[i]
		minHeapify(array, smallest)
	}
}

func buildMinHeapify(array []int) {
	// (n-1) to 0
	for i := (len(array) - 1) >> 1; i >= 0; i-- {
		minHeapify(array, i)
	}
}

func buildMaxHeapify(array []int) {
	// (n-1) to 0
	for i := (len(array) - 1) >> 1; i >= 0; i-- {
		maxHeapify(array, i)
	}
}

// O(n*logn)
func sort(array []int) {
	buildMaxHeapify(array)

	for i := len(array) - 1; i >= 1; i-- {
		array[0], array[i] = array[i], array[0]
		maxHeapify(array[:i], 0)
	}
}

// O(n*lgn): improve the performance
func minsort(array []int) {
	buildMinHeapify(array)

	for i := len(array) - 1; i >= 1; i-- {
		array[0], array[i] = array[i], array[0]
		minHeapify(array[:i], 0)
	}
}
