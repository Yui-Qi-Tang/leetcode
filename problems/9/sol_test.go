package solution

import (
	"testing"
)

func TestBasic(t *testing.T) {

	testcases := []int{1234567, 123456789, 1221, 12321, 11, 9, -1, 0, 11111}

	for _, testcase := range testcases {
		t.Log(testcase, isPalindrome3(testcase))
	}
}
