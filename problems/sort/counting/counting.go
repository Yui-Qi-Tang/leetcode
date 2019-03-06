package counting

// O(n * m) where n is the number of elements of input and m is the counts of element of input array
// BUG!!
func countingSortSlow(A, B []int, k int) {
	//Assume the max value = 5 of Array A
	C := make([]int, k) // default 0 in C
	for _, v := range A {
		C[v] = C[v] + 1
	}
	for i, v := range C {
		for j := 0; j < v; j++ {
			B = append(B, i)
		}
	}
}

func getMax(A []int) int {
	max := 0
	for _, v := range A {
		if v > max {
			max = v
		}
	}
	return max
}

func countingSort(A, B []int, k int) {
	C := make([]int, k) // default 0 in C
	for _, v := range A {
		C[v] = C[v] + 1
	}
	for i := 0; i < k; i++ {
		preIdx := i - 1
		if preIdx < 0 {
			continue
		}
		C[i] = C[i] + C[preIdx]
	}
	for i := len(A) - 1; i >= 0; i-- {
		B[C[A[i]]-1] = A[i]   // index is start with 0 -> the element position to B is  position - 1
		C[A[i]] = C[A[i]] - 1 // element counts - 1
	}
}
