package sol

/*
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkSol1-8         186785784                6.317 ns/op           0 B/op          0 allocs/op
BenchmarkSol2-8         158142132                7.601 ns/op           0 B/op          0 allocs/op
BenchmarkSol3-8         171271831                7.078 ns/op           0 B/op          0 allocs/op
*/

// findPeakElement returns the peak pos of the given array
// O(logn)
func findPeakElement(nums []int) int {
	return sol1(nums)
}

func sol1(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		m := mid(left, right)
		if nums[m] > nums[m+1] { // win right side
			right = m // find left part
		} else { // lose from rigth
			left = m + 1 // next round, left win its left-side
		}
	}
	return left
}

func sol2(nums []int) int {
	return peakfindingLeftFirst(nums, 0, len(nums)-1)
}

func sol3(nums []int) int {
	return peakfindingRightFirst(nums, 0, len(nums)-1)
}

func peakfindingLeftFirst(nums []int, i, j int) int {
	m := mid(i, j)
	if m > 0 && nums[m-1] > nums[m] { // left is greater than m
		return peakfindingLeftFirst(nums, i, m-1) // 0 to m-1
	} else if m < len(nums)-1 && nums[m] < nums[m+1] { // right is greater than m
		return peakfindingLeftFirst(nums, m+1, j)
	}
	// nums[m] >= nums[m-1] && nums[m] >= nums[m+1]
	return m
}

func peakfindingRightFirst(nums []int, i, j int) int {
	m := mid(i, j)
	if m < len(nums)-1 && nums[m] < nums[m+1] { // right is greater than m
		return peakfindingRightFirst(nums, m+1, j) // m+1 ... j
	} else if m > 0 && nums[m-1] > nums[m] { // left is greater than m
		return peakfindingRightFirst(nums, i, m-1) // i ... m-1
	}
	// nums[m] >= nums[m-1] && nums[m] >= nums[m+1]
	return m
}

func toPos(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func mid(x, y int) int {
	return (toPos(x) + toPos(y)) / 2
}
