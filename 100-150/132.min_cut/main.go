package main

import (
	"fmt"
	"strings"
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

func longestPalindrome(s string) (int, int) {
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
	return (2*maxI-(maxJ-1))/2, (maxJ-1)/2
}

