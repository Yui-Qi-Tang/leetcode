package quicksort

func sort(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)
		sort(nums, p, q-1)
		sort(nums, q+1, r)
	}
}

// TODO: try another way to pick the pivot
func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
