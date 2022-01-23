package sol

// 1143. Longest Common Subsequence

// Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.

// HINT: Must be considered the 'order' of comparsing!

// botton-up dp
// time complexity: O(m*n) where m is length of text1 and n is length of text2
// space: theta(m*n) where m is length of text1 + 1 and n is length of text2 + 1, '+1' is for the base case
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1] // add 1 with [i+1][j+1], becase the 'order', "it can not be added previous charator"
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1]) // compare the max of the previous record
			}
		}

	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
