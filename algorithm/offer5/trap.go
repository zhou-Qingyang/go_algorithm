func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	stack := make([]int, 0, n)
	ans := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			h := min(height[left], height[i]) - height[top]
			ans += width * h
		}
		stack = append(stack, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}