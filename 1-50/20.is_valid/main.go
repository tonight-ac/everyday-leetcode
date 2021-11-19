package main

// 所有case公用一个stack 节约内存 s最大长度10000 最多全放进来
var stack = make([]uint8, 10000)
// top表示栈顶部位置
var top int

// 很常见，用栈可以解决
func isValid(s string) bool {
	top = 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 入栈左号
			stack[top] = s[i]
			top++
		} else if top == 0 || s[i] - stack[top-1] > 2 {
			// 如果栈里没有元素说明右号多余，或者根据ASCII编码差距大于2，可知不匹配
			// https://www.matools.com/code-convert-ascii
			return false
		} else {
			// 匹配，出栈一个左号
			top--
		}
	}

	// 栈内是否还有未匹配的左号
	return top == 0
}