package main

import (
	"fmt"
	"math"
)

//对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
//
//给定一个 整数 n， 如果是完美数，返回 true，否则返回 false
//
// 
//
//示例 1：
//
//输入：num = 28
//输出：true
//解释：28 = 1 + 2 + 4 + 7 + 14
//1, 2, 4, 7, 和 14 是 28 的所有正因子。

func main() {
	fmt.Println(checkPerfectNumber(6))
}

func checkPerfectNumber(num int) bool {
	if num == 1 { return false }

	limit := math.Sqrt(float64(num))

	count := 0

	for i := float64(1); i <= limit; i++ {
		// 符合要求
		if num % int(i) == 0 {
			count += int(i)
			if i != 1 {
				count += num / int(i)
			}
		}
		// 提前结束
		if count > num {
			return false
		}
	}

	return count == num
}


