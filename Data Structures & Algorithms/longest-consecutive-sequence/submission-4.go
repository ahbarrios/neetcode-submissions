func longestConsecutive(nums []int) int {
	var max int
	nmap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		nmap[n] = struct{}{}
	}
	for _, n := range nums {
		if _, ok := nmap[n - 1]; ok {
			continue
		}
		a := n
		s := 1
		for _, ok := nmap[a + 1]; ok; {
			_, ok = nmap[a + 1]
			if ok {
				a++
				s++
			}
		}
		if s > max {
			max = s
		}
	}
	return max
}
