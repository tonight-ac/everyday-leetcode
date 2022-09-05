package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("ABC", 1))
}

// 1       9 2 * numRows - 2
// 2     8 1
// 3   7
// 4 6
// 5

// P   A   H   N
// A P L S I I G
// Y   I   R
func convert(s string, numRows int) string {
	// 排成一条直线的直接返回
	if numRows == 1 || len(s) <= numRows { return s }

	// 计算一个全周期的步长
	cycle := 2 * numRows - 2
	// 使用+拼接字符串消耗性能和内存
	var res strings.Builder

	// 处理首行
	if numRows > 0 {
		for idx := 0; idx < len(s); idx += cycle {
			res.WriteByte(s[idx])
		}
	}
	for row := 1; row < numRows - 1; row++ {
		// 计算需要跳步次数
		diff := cycle - 2 * row
		for idx := row; idx < len(s); {
			res.WriteByte(s[idx])
			idx += diff
			diff = cycle - diff
		}
	}
	// 处理尾行
	if numRows > 1 {
		for idx := numRows - 1; idx < len(s); idx += cycle {
			res.WriteByte(s[idx])
		}
	}

	return res.String()
}
