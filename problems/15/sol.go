package sol

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue // avoid to duplicated computing
		}
		target := nums[i]
		// 2 pointers
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := target + nums[low] + nums[high]
			if sum == 0 { // satify condition
				result = append(result, []int{target, nums[low], nums[high]}) // save result

				// try to find another pair to satisfy the target
				for low++; low < high && nums[low] == nums[low-1]; {
					low++ //skip the duplicated value
				}
				for high--; low < high && nums[high] == nums[high+1]; {
					high-- //skip the duplicated value
				}
			} else if sum < 0 {
				low++ // find the small one (HINT: nums is sorted)
			} else {
				high-- // find the bigger one (HINT: nums is sorted)
			}
		}
	}
	return result
}
