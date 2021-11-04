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
	fmt.Println(ss)
	// babad
	// #b#a#b#a#d#
	// 012345678910

	for i = 0; i < len(ss); i++ {
		for j = i + 1; j < len(ss); j++ {
			if 2 * i - j < 0 || ss[j] != ss[2 * i - j] {
				break
			}
		}
		fmt.Println(i, j)
		// 需要移动中心，先记下最大值
		if j - i > maxJ - maxI {
			maxJ, maxI = j, i
		}
	}
	fmt.Println(maxI, maxJ)

	return s[maxI/2:maxJ/2]
}
