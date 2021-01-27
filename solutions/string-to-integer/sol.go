// Package sti implements atoi which converts a string to an integer.
// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
//
// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
//
// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
//
// If no valid conversion could be performed, a zero value is returned.
//
// Note:
//
// Only the space character ' ' is considered as whitespace character.
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
package sti

import (
	"strconv"
	"strings"

	"regexp"
)

const max int = 1<<31 - 1
const min int = -1 << 31

func responseWithRange(x int) int {
	if x > max {
		return max
	} else if x < min {
		return min
	}
	return x
}

// myAtoi implements string to integer problem of leetcode
// but, the performance is very slow...
// Report from leetcode QQ
// Runtime: 8 ms, faster than 9.83% of Go online submissions for String to Integer (atoi).
// Memory Usage: 6.3 MB, less than 5.16% of Go online submissions for String to Integer (atoi).
// TODO: better performance
func myAtoi(str string) int {
	var validNum = regexp.MustCompile(`^[ ]*[\-\+]*[0-9]+`)
	data := validNum.FindString(str)
	if len(data) == 0 {
		return 0
	}

	var validOpt = regexp.MustCompile(`^[\-\+]{1}[0-9]+|^[0-9]+`)
	nStrWithOpt := strings.TrimSpace(data) // valid cases: '-', '+' otherwise invalid
	validString := validOpt.FindString(nStrWithOpt)

	nStr := strings.Trim(strings.TrimSpace(validString), "+")

	num, err := strconv.Atoi(nStr)
	if err != nil {
		if len(nStr) > 12 {
			if nStr[0] == '-' {
				return min
			}
			return max
		}
		//return 0
	}

	return responseWithRange(num)
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for String to Integer (atoi).


Bad code here lol
*/
func myAtoiBetter(s string) int {

	strim := strings.TrimSpace(s) // remove whitespace
	if len(strim) == 0 {
		return 0
	}

	sign := 1
	countSign := 0
	result := 0

	// check if firstChar of s without whitespace
	// expected chars := [0-9, +, -]
	firstChar := string(strim[0])
	if firstChar != "-" && // not equal "-" and
		firstChar != "+" && // not equal "+" and
		firstChar < "0" || // less than "0" ||
		firstChar > "9" { // greater than "9"
		// -> char does exist in expected chars
		return 0
	}

	for i, st := range strim {
		if string(st) >= "0" && string(st) <= "9" {
			result = result*10 + (int(st) - int('0'))
			// check next one is whether digit or not
			if i+1 < len(strim) {
				if int(strim[i+1]) < int('0') || int(strim[i+1]) > int('9') {
					return responseWithRange(result * sign)
				}
			}
		}

		if string(st) == "-" {
			if i+1 < len(strim) {
				if int(strim[i+1]) < int('0') || int(strim[i+1]) > int('9') {
					return 0
				}
			}
			sign = -1
			countSign++
		} else if string(st) == "+" {
			if i+1 < len(strim) {
				if int(strim[i+1]) < int('0') || int(strim[i+1]) > int('9') {
					return 0
				}
			}
			sign = 1
			countSign++
		}

		if countSign > 1 {
			return 0
		}

		// avoid the buffer overflow
		if result > max || result < min {
			return responseWithRange(result * sign)
		}
	}
	return responseWithRange(result * sign)
}
