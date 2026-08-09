func trap(height []int) int {
	var acc int

	maxLeft := make([]int, len(height))
	for i, max := 0, height[0]; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
		maxLeft[i] = max
	}

	maxRight := make([]int, len(height))
	for i, max := len(height) - 1, height[len(height)-1]; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		maxRight[i] = max
	}

	for i, h := range height {
		if a := min(maxLeft[i], maxRight[i]) - h; a > 0 {
			acc += a
		}
	}

	return acc
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
