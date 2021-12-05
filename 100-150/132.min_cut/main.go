package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCut("adabdcaebdcebdcacaaaadbbcadabcbeabaadcbcaaddebdbddcbdacdbbaedbdaaecabdceddccbdeeddccdaabbabbdedaaabcdadbdabeacbeadbaddcbaacdbabcccbaceedbcccedbeecbccaecadccbdbdccbcbaacccbddcccbaedbacdbcaccdcaadcbaebebcceabbdcdeaabdbabadeaaaaedbdbcebcbddebccacacddebecabccbbdcbecbaeedcdacdcbdbebbacddddaabaedabbaaabaddcdaadcccdeebcabacdadbaacdccbeceddeebbbdbaaaaabaeecccaebdeabddacbedededebdebabdbcbdcbadbeeceecdcdbbdcbdbeeebcdcabdeeacabdeaedebbcaacdadaecbccbededceceabdcabdeabbcdecdedadcaebaababeedcaacdbdacbccdbcece"))
}

func minCut(s string) int {
	n := len(s)
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
		for j := range g[i] {
			g[i][j] = true
		}
	}
	// 计算i～j是否为回文串
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
		}
	}

	f := make([]int, n)
	for i := range f {
		if g[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if g[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	return f[n-1]
}

//输入：
//"adabdcaebdcebdcacaaaadbbcadabcbeabaadcbcaaddebdbddcbdacdbbaedbdaaecabdceddccbdeeddccdaabbabbdedaaabcdadbdabeacbeadbaddcbaacdbabcccbaceedbcccedbeecbccaecadccbdbdccbcbaacccbddcccbaedbacdbcaccdcaadcbaebebcceabbdcdeaabdbabadeaaaaedbdbcebcbddebccacacddebecabccbbdcbecbaeedcdacdcbdbebbacddddaabaedabbaaabaddcdaadcccdeebcabacdadbaacdccbeceddeebbbdbaaaaabaeecccaebdeabddacbedededebdebabdbcbdcbadbeeceecdcdbbdcbdbeeebcdcabdeeacabdeaedebbcaacdadaecbccbededceceabdcabdeabbcdecdedadcaebaababeedcaacdbdacbccdbcece"
//输出：
//277
//预期结果：
//273
// 借用第5题
// 每次都寻找最长的回文字串，这种思路无法通过，但是对于这种巨长的case，我无法区分为什错
//func minCut(s string) int {
//	if len(s) <= 1 { return 0 }
//
//	i, j := longestPalindrome(s)
//	fmt.Println(s[i:j])
//
//	// 没有更多的子串了
//	// 最长就是1了，所以每个都需要分割
//	if j == i + 1 { return len(s) - 1 }
//
//	left, right := 0, 0
//	// 左侧还有，需要割一刀
//	if i != 0 {
//		left = minCut(s[:i]) + 1
//	}
//	// 右侧还有，需要割一刀
//	if j != len(s) {
//		right = minCut(s[j:]) + 1
//	}
//
//	return left + right
//}
//
//// 参考第5题
//func longestPalindrome(s string) (int, int) {
//	if len(s) == 0 { return 0, 0 }
//
//	maxR, maxL := 0, 0
//
//	// 单中心
//	for i := 0; i < len(s); i++ {
//		l, r := i, i
//		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
//			if s[l] != s[r] { break }
//		}
//		if maxR - maxL < r - l { maxR, maxL = r, l }
//	}
//
//	// 双中心
//	for i := 1; i < len(s); i++ {
//		l, r := i-1, i
//		if s[l] != s[r] { continue }
//		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
//			if s[l] != s[r] { break }
//		}
//		if maxR - maxL < r - l { maxR, maxL = r, l }
//	}
//
//	return maxL+1, maxR
//}
//
