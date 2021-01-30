package sol

/*
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
*/
var valueSymbol map[int]string = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",

	4:   "IV",
	9:   "IX",
	40:  "XL",
	90:  "XC",
	400: "CD",
	900: "CM",
}

func pow(base, times int) int {

	if base == 0 {
		return 0
	}

	if times == 0 {
		return 1
	}

	result := 1

	for i := 0; i < times; i++ {
		result *= base
	}

	return result
}

func intToRoman(num int) string {

	result := ""

	currentPosition := 0 // current pos of num
	for num > 0 {
		latestP := num % 10 // get latest pos of number
		value := latestP * pow(10, currentPosition)

		if symbol, exist := valueSymbol[value]; exist {
			result = symbol + result
		} else {
			result = toSymbol(value) + result
		}

		// push to next number
		currentPosition++
		num /= 10 // remove latest pos
	}

	return result
}

func toSymbol(v int) string {
	vcopy := v
	s := ""
	vmod := 0

	for vcopy > 0 {
		sym, exist := valueSymbol[vcopy]
		if exist {
			s = s + sym
			vcopy = vmod // just removed the diff and sync with the mod of v
			vmod = 0
		} else {
			vmod++
			vcopy--
		}

		if vcopy == 0 && vmod > 0 {
			vcopy = vmod
			vmod = 0
		}
	}
	return s
}
