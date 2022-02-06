package sol

// given an unsiged integer array 'nums', search the n is whether in nums
// return -1 if n is not in nums otherwise return the target as int

func searchLinear(nums []uint, n uint) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			return int(n)
		}
	}
	return -1
}

func searchDNC(nums []uint, n uint) int {
	return searchdnc(nums, n)
}

func searchdnc(nums []uint, n uint) int {

	if len(nums) == 0 {
		return -1
	}

	// base
	if len(nums) == 1 {
		if nums[0] == n {
			return int(n)
		} else {
			return -1
		}
	}

	// divide
	m := len(nums) / 2

	l := searchdnc(nums[:m], n)
	if l != -1 {
		return l
	}
	r := searchdnc(nums[m+1:], n)
	if r != -1 {
		return r
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
