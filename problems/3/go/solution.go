package solution

func lengthOfLongestSubstring(s string) int {
	repeatChar := make(map[rune]int) // index of rune of s
	result := 0 // max length of substring without repeat charator
	repeatNext := 0 // without repeat rune index

	for index, c := range s {
        if v, got := repeatChar[c]; got { // check repeat
            repeatNext = max(repeatNext, v+1) // 檢查repeatChar重複的元素索引;若元素有重複，會寫進repeatNext中，這個值是重複元素的下一位
		}
		repeatChar[c] = index 
		result = max(result, index - repeatNext + 1)
	}

	return result
}


func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}
/*
func main() {
    testCases := []string{
		"abcabcbb",
		"testtest",
		"pwwkew",
		"bbbbbbb",
		"",
		" ",
		"ab",
		"abcd",
		"aab",
		"asjrgapa",
	}

	for _, v := range testCases {
		fmt.Println("======>")
		fmt.Printf("Case %s: %d\n", v, lengthOfLongestSubstring(v))

	}
}*/