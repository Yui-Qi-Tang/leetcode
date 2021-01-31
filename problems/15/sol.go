package sol

import "fmt"

func threeSum(nums []int) [][]int {

	diff := make(map[int]int)
	result := make([][]int, 0)

	if len(nums) == 0 {
		return nil
	}

	for _, n := range nums {
		_, ok := diff[n]
		if !ok {
			diff[n] = 0 - n
		}
	}

	fmt.Println(diff)

	return result
}
