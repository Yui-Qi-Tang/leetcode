package main

import (
	"fmt"
)

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

func main() {
	var trianglehigh = 15
	for i := 0; i < trianglehigh; i++ {
		// Just add white space for view
		for k := trianglehigh - i; k > 0; k-- {
			fmt.Printf(" ")
		} // for
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", combination(i, j))
		}
		fmt.Printf("\n")
	}
}
