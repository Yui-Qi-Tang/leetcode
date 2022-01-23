package sol

func valid(i int) bool {
	if i < 0 || i > 10000 {
		return false
	}
	return true
}

// smallest if at left side and biggest is at right
// if left == right -> return 0
// if left < right -> ok
// if left > right -> no
func maxProfit(prices []int) int {

	if len(prices) == 0 || len(prices) > 100000 {
		return 0
	}

	max := 0
	var i, j int = 0, 0

	i = 0
	j = len(prices) - 1

	for i != j { // if they meet with each other which implice all elementes of prices are scan
		if !valid(prices[i]) || !valid(prices[j]) {
			return 0
		}

		if prices[j]-prices[i] > max {
			max = prices[j] - prices[i]
		}

		if prices[i] >= prices[j] {
			if prices[i] < prices[i+1] {
				j--
			} else {
				i++
			}
		} else { // prices[i] < prices[j]
			if prices[j] > prices[j-1] {
				i++
			} else {
				j--
			}
		}
	}

	return max
}

func maxProfit2(prices []int) int {

	if len(prices) == 0 || len(prices) > 100000 {
		return 0
	}

	max := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}

	return max
}

func maxProfit3(prices []int) int {

	if len(prices) == 0 || len(prices) > 100000 {
		return 0
	}

	l := len(prices) / 2
	r := len(prices) - 1

	lmin := 999999
	for i := 0; i < l; i++ {
		if prices[i] < lmin {
			lmin = prices[i]
		}
	}

	lmax := 0
	for i := l; i < r; i++ {
		if prices[i] > lmax {
			lmax = prices[i]
		}
	}

	if lmin >= lmax {
		return 0
	} else {
		return lmax - lmin
	}

}

func maxProfit4(prices []int) int {
	if len(prices) == 0 || len(prices) > 100000 {
		return 0
	}

	buy := prices[0]
	maxProfit := 0 // store the previous max profit

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buy)
		}
	}
	return maxProfit

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
