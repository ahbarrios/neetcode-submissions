func trap(height []int) int {
	var acc int 
	for l, r, maxL, maxR := 0, len(height) - 1, height[0], height[len(height)-1]; l <= r; {
		if maxL < maxR {
			maxL = max(maxL, height[l])
			acc += maxL - height[l]
			l++
		} else {
			maxR = max(maxR, height[r])
			acc += maxR - height[r]
			r--
		}
	}
	return acc
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
