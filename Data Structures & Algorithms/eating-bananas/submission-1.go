import "slices"

func minEatingSpeed(piles []int, h int) int {
	maxp := slices.Max(piles)

	hours := func (k int) int {
		var c int
		for _, p := range piles {
			c += (p + k - 1) / k
		}
		return c
	}

	var k int
	for d := maxp; d > 0; d = d/2 {
		for k + d <= maxp && hours(k + d) > h {
			k += d
		}
	}

	return k + 1
}
