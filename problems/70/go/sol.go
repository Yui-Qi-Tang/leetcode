package solution

//1 <= n <= 45

var steps map[int]int = make(map[int]int)

func climbStairs(n int) int {

	if n < 0 {
		return 0
	}

	if n == 1 || n == 0 {
		return 1 // just 1 step
	}

	if v, ok := steps[n]; ok {
		return v
	}

	steps[n] = climbStairs(n-1) + climbStairs(n-2)

	return steps[n]
}
