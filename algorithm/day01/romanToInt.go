package day01

//13
func RomanToInt(rom string) int {
	var value = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 1; i < len(rom); i++ {
		// 如果当前数字小于下一个数字，则减去当前数字
		if value[rom[i-1]] < value[rom[i]] {
			result -= value[rom[i-1]]
		} else { // 否则加上当前数字
			result += value[rom[i-1]]
		}
	}
	// 加上最后一个字符对应的整数值
	result += value[rom[len(rom)-1]]
	return result
}
