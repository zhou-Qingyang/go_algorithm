package day01

//20.有效的括号
func IsValid(s string) bool {
	stack := []byte{} // 定义一个栈
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if (s[i] == ')' && stack[len(stack)-1] != '(') || (s[i] == '}' && stack[len(stack)-1] != '{') || (s[i] == ']' && stack[len(stack)-1] != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
