package offer02

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '

}
