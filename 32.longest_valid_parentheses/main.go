package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("()(()"))
}
// TODO 未完成
func longestValidParentheses(s string) int {
	// 如果出现多余右括号，直接更新最大连续数字，后重新计数
	// 对于左括号，需要在最后结束的时候再统计是否匹配，所以需要把下标保存下来，
	// 方便直接计算长度
	left := 0
	leftIdxs := make([]int, 0)
	max, count := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // 左括号
			left++
			leftIdxs = append(leftIdxs, i) // 计算最新一个左括号下标
		} else { // 右括号
			if left > 0 { // 正常情况，左括号够用
				left--
				leftIdxs = leftIdxs[:len(leftIdxs)-1]
				count += 2
			} else if max < count { // 特殊情况，右括号多余，立刻计算长度
				max = count
				count = 0 // 开始下一轮计数
			}
		}
	}

	// 处理左括号
	if left != 0 && len(s) - 1 - leftIdxs[len(leftIdxs)-1] > max {
		max = len(s) - 1 - leftIdxs[len(leftIdxs)-1]
	} else if max < count { // 刚好匹配，最后更一下最大值
		max = count
	}

	return max
}