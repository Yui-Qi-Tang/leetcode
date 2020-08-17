// Package solution Reverse Integer of LeetCode.
// Problem:
//
// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
//     Input: 123
// 	Output: 321
//
// Example 2:
//     Input: -123
// 	Output: -321
//
// Example 3:
//     Input: 120
// 	Output: 21
//
// [Note]:
// 	Assume we are dealing with an environment whichcould only store integers within
// 	the 32-bit signed integer range: '[−2^31,  2^31 − 1]'.
// 	For the purpose of this problem, assume that your function returns 0 when the 'reversed' integer overflows.
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

	//  2020-08-17 檢查 32-bit signed integer 範圍，不只輸入要檢查，輸出的答案也要檢查(我忘記檢查輸出的範圍，導致錯了三次)。
	if outOfRange(result) {
		return 0
	}

	return result

}

// reverseBetter is an improvement of reverse function
func reverseBetter(x int) int {
	if outOfRange(x) {
		return 0
	}

	result := 0 // 紀錄反轉過後的數字

	for x != 0 {
		result = result*10 + x%10 // 利用餘數除法取得最小位數的值，並將之推前一位
		x = x / 10                // 去掉算過的位數
	}

	if outOfRange(result) {
		return 0
	}

	return result
}
