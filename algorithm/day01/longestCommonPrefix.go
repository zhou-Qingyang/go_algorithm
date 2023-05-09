package day01

//14
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs[0]) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}

	}
	// 返回基准字符串
	return strs[0]
}
