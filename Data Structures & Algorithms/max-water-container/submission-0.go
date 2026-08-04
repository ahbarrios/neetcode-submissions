func maxArea(heights []int) int {
	var max int
	for i, j := 0, len(heights) - 1; i < j; {
		a, b := heights[i], heights[j]

		mh := min(a, b)
		if s := mh * (j - i); s > max {
			max = s
		}

		if mh == a {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
