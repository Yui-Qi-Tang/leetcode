package sti

import (
	"testing"
)

// Test cases from leetcode:
//  Example 1:
//      Input: "42"
//      Output: 42
//
//  Example 2:
//      Input: "   -42"
//      Output: -42
//      Explanation: The first non-whitespace character is '-', which is the minus sign.
//                  Then take as many numerical digits as possible, which gets 42.
//  Example 3:
//      Input: "4193 with words"
//      Output: 4193
//      Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
//
//  Example 4:
//      Input: "words and 987"
//      Output: 0
//      Explanation: The first non-whitespace character is 'w', which is not a numerical
//                   digit or a +/- sign. Therefore no valid conversion could be performed.
//  Example 5:
//      Input: "-91283472332"
//      Output: -2147483648
//      Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
//                   Thefore INT_MIN (−231) is returned.

func TestBasic(t *testing.T) {

	testCases := []struct {
		input  string
		answer int
	}{
		{"    -4193 with words", -4193},
		{"    +4193 with words", 4193},
		{"    4193 with words", 4193},
		{"-4193", -4193},
		{"4193", 4193},
		{"error input", 0},
		{"42", 42},
		{"    -42", -42},
		{"-91283472332", -2147483648},
		{"words and 987", 0},
		{"+-2", 0},
		{"-+2", 0},
		{"20000000000000000000", 2147483647},
		{"-20000000000000000000", -2147483648},
		{"+-20000000000000000000", 0},
		{"  0000000000012345678", 12345678},
		{"  -0000000000012345678", -12345678},
		{"  +-0000000000012345678", 0},
		{"  -+0000000000012345678", 0},
		{"3.14159", 3},
		{" ", 0},
		{"9223372036854775808", 2147483647},
		{"  +  413", 0},
		{"+", 0},
		{" + ", 0},
	}

	for _, testcase := range testCases {
		result := myAtoi(testcase.input)
		if testcase.answer != result {
			t.Log("input:", testcase.input)
			t.Fatal("expected:", testcase.answer, "but got:", result)
		}
	}

	t.Log("... Passed")
}

func TestBetter(t *testing.T) {

	testCases := []struct {
		input  string
		answer int
	}{
		{"    -4193 with words", -4193},
		{"    +4193 with words", 4193},
		{"    4193 with words", 4193},
		{"-4193", -4193},
		{"4193", 4193},
		{"error input", 0},
		{"42", 42},
		{"    -42", -42},
		{"-91283472332", -2147483648},
		{"words and 987", 0},
		{"+-2", 0},
		{"-+2", 0},
		{"20000000000000000000", 2147483647},
		{"-20000000000000000000", -2147483648},
		{"+-20000000000000000000", 0},
		{"  0000000000012345678", 12345678},
		{"  -0000000000012345678", -12345678},
		{"  +-0000000000012345678", 0},
		{"  -+0000000000012345678", 0},
		{"3.14159", 3},
		{" ", 0},
		{"9223372036854775808", 2147483647},
		{"  +  413", 0},
		{"+", 0},
		{" + ", 0},
	}

	for _, testcase := range testCases {
		result := myAtoiBetter(testcase.input)
		if testcase.answer != result {
			t.Log("input:", testcase.input)
			t.Fatal("expected:", testcase.answer, "but got:", result)
		}
	}

	t.Log("... Passed")
}
