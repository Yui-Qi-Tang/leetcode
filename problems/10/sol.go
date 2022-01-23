package sol

/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).



Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*" => . or .. or ... or ........ => a(.)* -> ab(.)* -> ab
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:

Input: s = "mississippi", p = "mis*is*p*."
Output: false


Constraints:

0 <= s.length <= 20
0 <= p.length <= 30
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.


// regexp.MustCompile("^" + p + "$").MatchString(s) //lol
*/

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	} else if p[len(p)-1] == '*' {
		return isMatch(s, p[:len(p)-2]) || // check preceding element
			// last 2 of p is '.' ->  any single char(XXX(.*)XXX) where XXX is other chars || last one of s is equal to last 2 of p
			len(s) > 0 && (p[len(p)-2] == '.' || s[len(s)-1] == p[len(p)-2]) &&
				isMatch(s[:len(s)-1], p) // for what?
	} else if len(s) > 0 && (p[len(p)-1] == '.' || s[len(s)-1] == p[len(p)-1]) { // check if last one of p is '.' or last one of s is equal to last one of p
		return isMatch(s[:len(s)-1], p[:len(p)-1])
	}
	return false
}

func solV2(s, p string) bool {

	// dp[s[i]][p[j]] is a comparsion of s[i] and p[j]
	dp := make([][]bool, len(s)+1) // +1 for s is empty

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(p)+1) // +1 for p is empty
	}

	dp[0][0] = true // |p|==0 and |s|==0

	return false
}
