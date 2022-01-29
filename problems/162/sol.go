package sol

import "math"

var min int = -math.MaxInt

// O(n)
func findPeakElement(nums []int) int {

	i := 0    // point of nums
	peak := 0 // index of peak

	for i < len(nums) {

		p := 0
		if i-1 < 0 {
			p = min
		} else {
			p = nums[i-1]
		}

		c := nums[i]

		n := 0
		if i+1 > len(nums)-1 {
			n = min
		} else {
			n = nums[i+1]
		}

		if isPeak(p, c, n) {
			return i
		} else {
			i++
		}
	}
	return peak
}

// O(logn)
func findPeakElementv2(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		m := left + (right-left)/2
		if nums[m] > nums[m+1] {
			right = m
		} else {
			left = m + 1
		}
	}

	return left
}

func peak1d(nums []int, i, j int) int {
	m := len(nums) / 2

	p := min // point to m-1
	if m-1 >= 0 {
		p = nums[m-1]
	}

	c := nums[m]

	n := min
	if m+1 < len(nums) {
		n = nums[m+1]
	}

	if p > c {
		return peak1d(nums, i, m)
	} else if n > c {
		return peak1d(nums, m+1, j)
	}

	return m

}

func isPeak(p, c, n int) bool {
	if c > p && c > n {
		return true
	}
	return false
}
