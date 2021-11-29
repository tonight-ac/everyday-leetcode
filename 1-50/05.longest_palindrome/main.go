package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	maxI, maxJ := 0, 0

	ss := "#" + strings.Join(strings.Split(s, ""), "#") + "#"

	i, j := 0, 0

	for i = 0; i < len(ss); i++ {
		for j = i + 1; j < len(ss); j++ {
			if 2 * i - j < 0 || ss[j] != ss[2 * i - j] {
				break
			}
		}
		// 需要移动中心，先记下最大值
		if j - i > maxJ - maxI {
			maxJ, maxI = j, i
		}
	}

	// 2*maxI-(maxJ-1) 左边界
	// 2:(maxJ-1) 右边界
	return s[(2*maxI-(maxJ-1))/2:(maxJ-1)/2]
}
