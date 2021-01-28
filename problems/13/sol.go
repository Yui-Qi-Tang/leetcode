package sol

var valueSymbol map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func sub(a, b byte) (int, bool) {
	if a == 'I' && b == 'V' || a == 'I' && b == 'X' ||
		a == 'X' && b == 'L' || a == 'X' && b == 'C' ||
		a == 'C' && b == 'D' || a == 'C' && b == 'M' {
		return valueSymbol[b] - valueSymbol[a], true
	}
	return 0, false
}

func romanToInt(s string) int {

	slen := len(s)
	result := 0

	i := 0
	for i < slen {
		j := i + 1
		if j < slen {
			subv, ok := sub(s[i], s[j])
			if ok {
				result += subv
				i = j + 1 // skip to the next of j
				continue
			} else {
				result += valueSymbol[s[i]]
			}
		} else {
			result += valueSymbol[s[i]]
		}
		i++

	}

	return result
}
