func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r - l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func searchRange(nums []int, l, r, t int) int {
	m := l + (r - l) / 2
	if l > r {
		return -1
	}
	if nums[m] == t {
		return m
	}
	if t < nums[m] {
		return searchRange(nums, l, m, t)
	}
	return searchRange(nums, m + 1, r, t)
}
