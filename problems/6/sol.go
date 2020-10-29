package solution

/*
Input: s = "PAYPALISHIRING", numRows = 4

            ^^^^^^r0^^^^^^r1...

observe: (numRows =4)

col     0  12345  6  789AB  C
0   => 'P' AYPAL 'I' SHIRI 'N' G -> i+6=next

1   =>  P 'A' YPA 'L' I 'S' HIR 'I' N 'G'

0 => PAYP
1 => PALI
2 => ISHI
3 => IRIN
4 => NG**

Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I


row 0 index: 2 * numRows -2
numRows - 1 index : (2 * numRows - 2) + numRows - 1
row i index : k=(2 * numRows -2) + i and (k+1) (2*numRows-2) -i ??????
*/

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	runes := []rune(s)
	n := len(runes)
	result := ""

	rowCycke := (2*numRows - 2)

	// i controls row
	// j controls each element of row with indexing
	for i := 0; i < numRows; i++ {
		for j := 0; i+j < n; j = j + rowCycke {
			result += string(runes[j+i]) // append first char of row
			if i != 0 && i != numRows-1 && j+rowCycke-i < n {
				result += string(runes[j+rowCycke-i])
			}
		}
	}

	return result
}
