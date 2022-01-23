package sol

func containsDuplicate(nums []int) bool {
	meno := make(map[int]bool)

	for _, num := range nums {
		if _, dup := meno[num]; dup {
			return true
		}
		meno[num] = true
	}
	return false
}
