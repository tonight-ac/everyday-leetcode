package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

//示例 1：
//
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//示例 2：
//
//输入：s = "3[a2[c]]"
//输出："accaccacc"
//示例 3：
//
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//示例 4：
//
//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"

func main() {
	fmt.Println(decodeString("100[leetcode]"))
}

func decodeString(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			left := i
			// 向前寻找重复次数
			count := 0
			for j := left - 1; j >= 0 && unicode.IsDigit(rune(s[j])); j-- {
				// 处理多位数字
				count += int(s[j] - '0') * int(math.Pow10(left - 1 - j))
			}

			// 向后寻找右括号
			right := 0
			for pair, j := 1, left + 1; j < len(s); j++ {
				// pair处理多括号嵌套，模拟一个栈
				if s[j] == '[' {
					pair++
				} else if s[j] == ']' {
					if pair == 1 {
						right = j
						break
					}
					pair--
				}
			}

			// 递归处理
			b.WriteString(strings.Repeat(decodeString(s[left:right]), count))

			// 跳过
			i = right
		} else if !unicode.IsDigit(rune(s[i])) {
			b.WriteByte(s[i])
		}
	}

	return b.String()
}
