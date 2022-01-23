package sol

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).



Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.
Example 3:

Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
Example 4:

Input: s = "adceb", p = "*a*b"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
Example 5:

Input: s = "acdcb", p = "a*c?b"
Output: false
*/

func isMatch(s string, p string) bool {

	// mem[i][j] means isMatch(s[:i], p[:j])
	mem := make([][]bool, len(s)+1)
	for i := range mem {
		mem[i] = make([]bool, len(p)+1)
	}

	// init bound, mem[n][0] is false while n > 0
	mem[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			mem[0][j] = mem[0][j-1]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				mem[i][j] = mem[i][j-1] || mem[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				mem[i][j] = mem[i-1][j-1]
			}
		}
	}

	return mem[len(s)][len(p)]
}
