package pascaltriangle

/* hierarchy approach will lead buffer overflow on data type... (consider 100!)*/

import "fmt"

var mp map[int]int = make(map[int]int)

// c(n, r) = n! / r! * (n-r)!
func combination(n, r int) int {
	return hierarchy(n) / (hierarchy(r) * hierarchy(n-r))
}

// x!
func hierarchy(num int) int {
	if num < 0 {
		panic("num must be >= 0")
	}
	if num == 1 || num == 0 {
		return 1
	}
	return num * hierarchy(num-1)
}

// c(n, r) = n! / r! * (n-r)!
func combinationDP(n, r int) int {
	return hierarchyDP(n) / (hierarchyDP(r) * hierarchyDP(n-r))
}

func hierarchyDP(num int) int {
	if num < 0 {
		panic("num must be >= 0")
	}

	if num == 1 || num == 0 {
		return 1
	}

	if v, ok := mp[num]; ok {
		return v
	}

	n := num * hierarchyDP(num-1)
	mp[num] = n

	return n

}

func generate(numRows int) [][]int {
	if numRows < 0 {
		return nil
	}
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, 0)
		for j := 0; j < i+1; j++ {
			r := combinationDP(i, j)

			if r < 0 {
				fmt.Println(i, j)
			}
			result[i] = append(result[i], combinationDP(i, j))
		}
	}

	return result
}

// ref: https://ljun20160606.github.io/leetcode/algorithms/118/
func generate2(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, helper(res[i-1]))
	}

	return res
}

func helper(p []int) []int {
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i := 0; i < len(res)-1; i++ {
		res[i] += res[i+1]
	}

	return res
}
