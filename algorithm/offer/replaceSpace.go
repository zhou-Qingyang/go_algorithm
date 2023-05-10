package offer

import "bytes"

// 05
func replaceSpace1(s string) string {
	count := 0
	for _, c := range s {
		if c == ' ' {
			count++
		}
	}
	arr := make([]byte, len(s)+count*2)
	i, j := len(s)-1, len(arr)-1
	for i >= 0 {
		if s[i] == ' ' {
			arr[j] = '0'
			arr[j-1] = '2'
			arr[j-2] = '%'
			j -= 3
		} else {
			arr[j] = s[i]
			j--
		}
		i--
	}
	return string(arr)

}

// 05
func replaceSpace2(s string) string {
	var buffer bytes.Buffer
	// 遍历原字符串，如果当前字符是空格，则在 buffer 中添加%20；否则添加原字符。
	for _, c := range s {
		if c == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteRune(c)
		}
	}
	// 返回 buffer 转成的字符串
	return buffer.String()
}
