package offer02

import "strconv"

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	digit, base, num := 1, 9, 1 // digit 表示数字的位数，base 表示当前位数的数字总数，num 表示第一个数字
	for n-base*digit > 0 {
		n -= base * digit
		digit++
		base *= 10
		num *= 10
	}
	// 计算目标数字所在的原始数字 num 和该数字的位置 pos
	pos := (n - 1) % digit
	num += (n - 1) / digit
	// 将 num 转换成字符串，并取出目标数字
	numStr := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(numStr[pos]))
	return res
}
