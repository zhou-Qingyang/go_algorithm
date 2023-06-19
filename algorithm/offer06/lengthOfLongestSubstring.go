package offer06

//3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	lasPos := make(map[byte]int)
	left, res := 0, 0

	for right := 0; right < len(s); right++ {
		if pos, ok := lasPos[s[right]]; ok && pos <= left {
			left = pos + 1
		}
		lasPos[s[right]] = right
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
