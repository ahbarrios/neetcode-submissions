func largestRectangleArea(heights []int) int {
	var max int
	n := len(heights)
	for i, h := range heights {
		s := h
		mh := h
		for j := i+1; j < n; j++ {
			mh = min(mh, heights[j])
			if ss := mh * (j - i + 1); ss > s {
				s = ss
			}
		}
		if max < s {
			max = s
		}
	}
	return max
}

// helpers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// stack basic methods to handle on slice
func top(st []int) int {
	return st[len(st)-1]
}
func pop(st []int) ([]int, int) {
	v := top(st)
	return st[:len(st)-1], v
}
func push(st []int, v int) []int  {
	return append(st, v)
}
