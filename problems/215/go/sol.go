package sol

// heap approach

func left(i int) int {
	return i << 1
}

func right(i int) int {
	return i<<1 | 0x1
}

// find the max. value and let it as root
func maxHeapify(nums []int, i int) {
	l := left(i)
	r := right(i)

	largest := i
	if l < len(nums) && nums[l] > nums[i] {
		largest = l
	}

	if r < len(nums) && nums[r] > nums[largest] {
		largest = r
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest)
	}
}

func findKthLargest(nums []int, k int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	for i := len(nums) >> 1; i >= 0; i-- {
		maxHeapify(nums, i)
	}

	if k == 1 {
		return nums[0]
	}

	for i := len(nums) - 1; i > len(nums)-k; i-- {

		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums[:i], 0)

	}

	return nums[0]
}
