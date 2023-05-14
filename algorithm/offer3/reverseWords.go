package offer3

import "strings"

func reverseWords(s string) string {
	// 将整个句子翻转过来
	s = reverse(s)
	res := []byte{}
	start, end := 0, 0
	for end <= len(s) {
		if end == len(s) || s[end] == ' ' {
			// 翻转单词
			res = append(res, reverse(s[start:end])...)
			if end < len(s)-2 && s[end] == ' ' && s[end+1] == ' ' {
				start = end + 1
				end++
				continue
			}
			if end < len(s) {
				res = append(res, ' ')
			}
			start = end + 1
		}
		end++
	}
	return string(res)
}

func reverse(s string) string {
	s = strings.TrimSpace(s) // 去除左右两侧空格
	b := []byte(s)
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
