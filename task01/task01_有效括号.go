package task01

// 判断字符串是否为有效括号组合
func IsValid(s string) bool {
	stack := make([]rune, 0)
	runeMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != runeMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
