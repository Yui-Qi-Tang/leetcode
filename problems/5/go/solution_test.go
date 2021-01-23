package solution

import (
	"testing"
)

/*
Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
Example 3:

Input: s = "a"
Output: "a"
Example 4:

Input: s = "ac"
Output: "a"
*/

func TestSolution(t *testing.T) {
	testcases := []struct {
		input string
		want  string
	}{
		{input: "babad", want: "bab"},
		{input: "cbbd", want: "bb"},
		{input: "a", want: "a"},
		{input: "ac", want: "a"},
		{input: "baBaD0", want: "aBa"},
		{input: "a01010a1010101", want: "a01010a"},
		{input: "a01010a1*010101", want: ""},
		{input: "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", want: ""},
	}

	for _, tc := range testcases {
		if longestPalindrome(tc.input) != tc.want {
			t.Fatal("it should be", tc.want, "but got", longestPalindrome(tc.input))
		}
	}

	t.Log("passed")
}
