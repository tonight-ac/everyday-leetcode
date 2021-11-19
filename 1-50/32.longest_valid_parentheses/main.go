package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("()))()"))
}

func longestValidParentheses(s string) int {
	// 如果出现多余右括号，直接更新最大连续数字，后重新计数
	// 对于左括号，需要在最后结束的时候再统计是否匹配，所以需要把下标保存下来，
	// 方便直接计算长度

	max := 0
	stack := make([]int, 0)
	// 添加一个-1作为栈底
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // 左括号无脑加入
			stack = append(stack, i)
		} else {
			// 遇到右括号先出栈一个左括号
			if len(stack) > 0 { stack = stack[:len(stack)-1] }
			if len(stack) == 0 {
				// 出栈之后栈为空，则说明当前右括号不合法，入栈作为栈底
				stack = append(stack, i)
			} else {
				// 当前右括号匹配上了，直接计算最大值
				if max < i - stack[len(stack)-1] {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}

	return max
}