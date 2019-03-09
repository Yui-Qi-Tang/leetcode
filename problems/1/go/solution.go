package twosum2

func solution(target int, nums []int) []int {
	diffRecords := make(map[int]int)
	for i, v := range nums {
		d := target - v
		if _, ok := diffRecords[d]; ok {
			return []int{diffRecords[d], i}
		}
		diffRecords[v] = i
	}
	return nil
}
