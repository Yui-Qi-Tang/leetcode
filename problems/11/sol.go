package sol

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {

	hLen := len(height)
	max := 0
	for i, h := range height {
		for j := i + 1; j < hLen; j++ {
			area := min(h, height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}

	return max
}

// implemets two pointer
func maxAreaTwoPinter(height []int) int {

	hLen := len(height)
	left := 0
	right := hLen - 1
	m := 0
	for left < right {
		hl := height[left]
		hr := height[right]
		m = max(m, min(hl, hr)*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return m
}
