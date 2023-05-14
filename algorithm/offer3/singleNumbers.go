package offer3

func singleNumbers(nums []int) []int {
	// 计算所有数字的异或结果
	xorRes := 0
	for _, num := range nums {
		xorRes ^= num
	}

	// 找到两个数字不同的二进制位
	mask := 1
	for (mask & xorRes) == 0 {
		mask <<= 1
	}

	// 根据该二进制位的值，将数字分为两组
	a, b := 0, 0
	for _, num := range nums {
		if num&mask != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
