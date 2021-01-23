package solution

/*
   Problem: Given a string s, find the longest palindromic substring in s.
   Constraints:
       1 <= s.length <= 1000
       s consist of only digits and English letters (lower-case and/or upper-case),

*/
// madam, adda, babab, bb

// s consist of only digits and English letters (lower-case and/or upper-case),
func isvalid(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

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

	if !isvalid(s) {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if len(s) > 1000 {
		return ""
	}

	result := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if isPalindronmic(s[i:j]) && len(s[i:j]) > len(result) {
				result = s[i:j]
			}
		}
	}

	return result
}
