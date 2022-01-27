package sol

func maxProduct(nums []int) int {

	result, currMin, currMax := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		v := nums[i]

		if v < 0 { // +,- ; -, -
			// nagtive integer product with small integer which is greater than big integer
			// for example
			// v is -1 and currMin is 1, currMax is 2
			// --> v * 1 > v * 2
			// v is -1 and currMin is -100; currMax is -1
			// --> v * -100 > v * -1
			// so, if v is negative integer we, must swap the 'max' and 'min'
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
