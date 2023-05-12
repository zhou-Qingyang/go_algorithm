package offer02

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	res := 0
	m := make(map[byte]int, 0)
	for right < len(s) {
		if _, ok := m[s[right]]; ok && m[s[right]] >= left {
			left = m[s[right]] + 1
		}
		m[s[right]] = right
		res = max(res, right-left+1)
		right++
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
