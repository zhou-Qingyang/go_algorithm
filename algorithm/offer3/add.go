package offer3

func add(a int, b int) int {
	for b != 0 {
		// 计算不考虑进位时的和
		sum := a ^ b
		// 计算进位的情况
		carry := (a & b) << 1
		// 将两者相加，得到最终的和
		a, b = sum, carry
	}
	return a
}
