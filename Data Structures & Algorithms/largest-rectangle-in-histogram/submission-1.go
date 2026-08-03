func largestRectangleArea(heights []int) int {
	var max int
	heights = append(heights, 0)
	st := make([]int, 0, len(heights))
	id := make([]int, 0, len(heights))
	for i, h := range heights {
		ii := i
		for j := top(st); j >= 0 && heights[j] > h; j = top(st) {
			jj := j
			st, j = pop(st)
			id, jj = pop(id)
			if s := heights[j] * (i - jj); s > max {
				max = s
			}
			ii = jj
		}
		st = push(st, i)
		id = push(id, ii)
	}
	return max
}

// stack basic methods to handle on slice
func top(st []int) int {
	if len(st) > 0 {
		return st[len(st)-1]
	}
	return -1
}
func pop(st []int) ([]int, int) {
	v := top(st)
	if len(st) > 0 {
		return st[:len(st)-1], v
	}
	return st, v
}
func push(st []int, v int) []int  {
	return append(st, v)
}
