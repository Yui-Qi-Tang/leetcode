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

// bad performance, because I check the reverse of each substring
// n * (n - 1) * n => O(n^3)
func badLongestPalindrome(s string) string {

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

// check itself and next field
// n * 2n => O(n^2)
// LaR, LaaR where L and R statisfy palindrome, and
// base cases is
// a  is a single char -> statify palindrome
// aa is two single char and there are equal -> sataisfy palindrom
// if s == basecase -> head will be back to previous one and tail goes to next one (add the checking range)
// -> let the checking to L and R
// L <- a/aa -> R
// L is substring that is needed to satisfy above rule
// R is substring that is needed to satisfy above rule
func betterLongestPalindrome(s string) string {

	if !isvalid(s) {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if len(s) > 1000 {
		return ""
	}

	max := 0
	min := 0
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		prev := i
		tail := i
		for prev >= 0 && tail < sLen && s[prev] == s[tail] {
			prev--
			tail++
		}

		if tail-prev-1 > max-min {
			max = tail
			min = prev + 1
		}

		prev = i
		tail = i + 1
		for prev >= 0 && tail < sLen && s[prev] == s[tail] {
			prev--
			tail++
		}

		if tail-prev-1 > max-min {
			max = tail
			min = prev + 1
		}

	}

	return s[min:max]
}

func simple(s string) string {

	min, max := 0, 0
	sLen := len(s)
	for i := 0; i < len(s); i++ {
		head, tail := i, i

		for head >= 0 && tail < sLen && s[head] == s[tail] {
			head--
			tail++
		}

		if tail-head-1 > max-min {
			max = tail
			min = head + 1
		}

		head, tail = i, i+1

		for head >= 0 && tail < sLen && s[head] == s[tail] {
			head--
			tail++
		}

		if tail-head-1 > max-min {
			max = tail
			min = head + 1
		}

	}
	return s[min:max]
}
