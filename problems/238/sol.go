package sol

func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))

	for i := range nums {
		product := 1
		for j, ev := range nums {
			if i != j {
				product *= ev
			}
		}
		result = append(result, product)
	}

	return result
}

// O(n)
func productExceptSelf2(nums []int) []int {
	result := make([]int, 0, len(nums))
	zero := make(map[int]bool)
	total := 1

	for i, v := range nums {
		if v != 0 {
			total *= v
		} else {
			zero[i] = true
		}
	}

	for i, v := range nums {
		if len(zero) >= 2 {
			result = append(result, 0)
		} else if len(zero) > 0 {
			_, iszero := zero[i]
			if !iszero {
				result = append(result, 0)
			} else {
				result = append(result, total)
			}
		} else {
			result = append(result, total/v)
		}
	}

	return result
}
