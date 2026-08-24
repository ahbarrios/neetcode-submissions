func maxSlidingWindow(nums []int, k int) []int {
	st := make([]int, 0, len(nums))
	for i := 0; i < k; i++ {
		for !empty(st) && nums[top(st)] < nums[i] {
			st, _ = pop(st)
		}
		st = append(st, i)
	}
	maxs := make([]int, 0, len(nums))
	maxs = append(maxs, nums[front(st)])
    for r, l := 1, k; l < len(nums); l++ {
		for !empty(st) && front(st) < r {
			st, _ = fpop(st)
		}
		for !empty(st) && nums[top(st)] < nums[l] {
			st, _ = pop(st)
		}
		st = append(st, l)
		maxs = append(maxs, nums[front(st)])
		r++
    } 

    return maxs
}

func empty(st []int) bool {
	return len(st) == 0
}

func top(st []int) int {
	return st[len(st)-1]
}

func pop(st []int) ([]int, int) {
	n := top(st)
	return st[:len(st)-1], n
}

func front(st []int) int {
	return st[0]
}

func fpop(st []int) ([]int, int) {
	n := front(st)
	return st[1:], n
}