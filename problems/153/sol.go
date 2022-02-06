package sol

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.
*/

func findMin(nums []int) int {

	size := len(nums)
	// return -1 no elements in nums
	if size == 0 {
		return -1
	}

	if size == 1 {
		return nums[0]
	}

	left, right := 0, size-1

	for left < right {

		// the array is sorted
		if nums[left] <= nums[right] {
			return nums[left]
		}

		m := (left + right) / 2
		if nums[left] <= nums[m] {
			left = m + 1
		} else {
			right = m
		}
	}

	return nums[left]
}
