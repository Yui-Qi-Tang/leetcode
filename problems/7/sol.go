package solution

import (
	"strconv"
)

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func outOfRange(x int) bool {
	if x > 1<<31-1 || x < -1<<31 {
		return true
	}
	return false
}

// reverse solution here
func reverse(x int) int {

	if outOfRange(x) {
		return 0
	}

	xCopy := x

	if x < 0 {
		xCopy = xCopy * -1
	}

	result, err := strconv.Atoi(reverseString(strconv.Itoa(xCopy)))

	if err != nil {
		return 0
	}

	if x < 0 {
		result = result * -1
	}

	if outOfRange(result) {
		return 0
	}

	return result

}
