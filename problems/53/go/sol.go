package sol

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	i := 0
	maxNum := nums[0]

	partialTol := 0

	for i < len(nums) {
		partialTol += nums[i]

		if partialTol > maxNum {
			maxNum = partialTol
		}



		if partialTol < 0 {
			partialTol = 0
		}

		i++
	}

	return maxNum
}
