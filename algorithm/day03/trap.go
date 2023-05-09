package day03

//42
func Trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	leftMax, rightMax := height[0], height[n-1]
	res := 0
	for left <= right {
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
			if height[left-1] > leftMax {
				leftMax = height[left-1]
			}
		} else {
			res += rightMax - height[right]
			right--
			if height[right+1] > rightMax {
				rightMax = height[right+1]
			}
		}
	}
	return res
}
