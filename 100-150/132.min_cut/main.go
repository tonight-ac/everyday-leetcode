package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCut("ab"))
}

//很接近了
//输入：
//"apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"
//输出：
//453
//预期结果：
//452
// 借用第5题
func minCut(s string) int {
	if len(s) == 0 { return 0 }

	i, j := longestPalindrome(s)

	// 没有更多的子串了
	// 最长就是1了，所以每个都需要分割
	if j == i + 1 { return len(s) - 1 }

	res := 0
	left, right := 0, 0
	// 左侧还有，需要割一刀
	if i != 0 {
		left = minCut(s[:i])
		res++
	}
	// 右侧还有需要割一刀
	if j != len(s) {
		right = minCut(s[j:])
		res++
	}

	return res + left + right
}

// 参考第5题
func longestPalindrome(s string) (int, int) {
	if len(s) == 0 { return 0, 0 }

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

	return maxL+1, maxR
}

