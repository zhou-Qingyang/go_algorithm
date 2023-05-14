package offer3

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	q := make([]int, 0)

	for i := range nums {
		// 如果队列不为空且队头元素已经不在当前滑动窗口中了，将其出队
		if len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}

		// 如果当前元素大于队列中所有元素，将队列清空
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		// 将 i 入队
		q = append(q, i)

		// 如果队列长度大于等于 k，说明当前滑动窗口已经完整，队头元素即为当前窗口最大值
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}

	return res
}
