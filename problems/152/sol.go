package sol

func maxProduct(nums []int) int {

	result, currMin, currMax := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		v := nums[i]

		if v < 0 { // +,- ; -, -
			currMax, currMin = currMin, currMax
		}

		currMax = max(v*currMax, v)
		currMin = min(v*currMin, v)
		result = max(result, currMax)

	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
