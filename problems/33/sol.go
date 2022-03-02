package sol

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if nums[0] == target {
		return 0
	}

	size := len(nums) - 1
	i, j := 0, size

	for i <= j {
		m := (i + j) / 2

		if nums[m] == target {
			return m
		}

		// sorted
		if nums[i] <= nums[m] {
			if nums[i] <= target && target <= nums[m] {
				j = m - 1
			} else {
				i = m + 1
			}
		} else {
			if nums[m] <= target && target <= nums[j] {
				i = m + 1
			} else {
				j = m - 1
			}
		}

	}

	return -1

}
