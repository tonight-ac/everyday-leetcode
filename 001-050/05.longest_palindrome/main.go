package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 { return "" }

	maxR, maxL := 0, 0

	// 单中心
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] != s[r] { break }
		}
		if maxR - maxL < r - l { maxR, maxL = r, l }
	}

	// 双中心
	for i := 1; i < len(s); i++ {
		l, r := i-1, i
		if s[l] != s[r] { continue }
		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] != s[r] { break }
		}
		if maxR - maxL < r - l { maxR, maxL = r, l }
	}

	return s[maxL+1:maxR]
}

//func longestPalindrome(s string) string {
//	maxI, maxJ := 0, 0
//
//	ss := "#" + strings.Join(strings.Split(s, ""), "#") + "#"
//
//	i, j := 0, 0
//
//	for i = 0; i < len(ss); i++ {
//		for j = i + 1; j < len(ss); j++ {
//			if 2 * i - j < 0 || ss[j] != ss[2 * i - j] {
//				break
//			}
//		}
//		// 需要移动中心，先记下最大值
//		if j - i > maxJ - maxI {
//			maxJ, maxI = j, i
//		}
//	}
//
//	// 2*maxI-(maxJ-1) 左边界
//	// 2:(maxJ-1) 右边界
//	return s[(2*maxI-(maxJ-1))/2:(maxJ-1)/2]
//}
