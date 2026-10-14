func findMin(nums []int) int {
	n := len(nums)
	lo := -1
	hi := n-1
	for dif := hi - lo; dif > 0; dif = dif / 2 {
		for lo + dif <= hi && nums[lo + dif] >= nums[0] {
			lo += dif
		}
	}
	// lo := 0
	// for lo < hi {
	// 	mid := lo + (hi - lo) / 2
	// 	fmt.Println(lo, hi, mid)
	// 	if nums[mid] >= nums[0] {
	// 		lo = mid + 1
	// 	} else {
	// 		hi = mid
	// 	}
	// }
	return nums[(lo+1)%n]
}
