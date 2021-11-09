package main

import "fmt"

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
	res := ""

	for row := 0; row < numRows; row++ {
		// 计算需要跳步次数
		diff := cycle - 2 * row
		if diff == 0 {
			diff = cycle
		}
		for idx := row; idx < len(s); {
			res += string(s[idx])
			idx += diff
			// 跳步等于全周期步长做减法会出现0，导致原地踏步一次
			if diff != cycle {
				diff = cycle - diff
			}
		}
	}

	return res
}
