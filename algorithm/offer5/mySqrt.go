package offer5

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	for left <= right {
		mid := (left + right) / 2
		if mid > x/mid {
			right = mid - 1
		} else if (mid + 1) > x/(mid+1) {
			return mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
