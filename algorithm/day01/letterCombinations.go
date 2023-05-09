package day01

//17
func LetterCombinations(digits string) []string {
	if digits == "" { // 如果给定数字字符串为空，则返回空切片
		return []string{}
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	result := []string{""}
	for _, v := range digits {
		letters := m[byte(v)]
		nextResult := []string{} // 初始化下一轮遍历的结果切片
		for _, l := range result {
			for _, r := range letters {
				nextResult = append(nextResult, l+string(r))
			}
		}
		result = nextResult
	}
	return result // 返回最终结果切片
}
