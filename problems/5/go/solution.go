package solution

/*
   Problem: Given a string s, find the longest palindromic substring in s.
            You may assume that the maximum length of s is 1000.
*/
// madam, adda, babab, bb

func reverse(s string) string {
	var result string

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func isPalindronmic(s string) bool {
	sReverse := reverse(s)

	for i, v := range s {
		if v != rune(sReverse[i]) {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {

	return ""
}
