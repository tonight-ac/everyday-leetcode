package main

import "strings"

func longestPalindrome(s string) string {
	res := ""
	det := 0
	m := make(map[uint8]map[int]bool)

	ss := strings.Split(s, "")

	s = "#" + strings.Join(ss, "#") + "#"

	for i := 0; i < len(s); i++ {
		v := s[i]
		m[v][i] = true

		// 如果对称位找到，则弹针向前移动1
		if idx, ok := m[v][det-i]; ok {
			det++
		} else { // 没有找到则需要更新中间值位置，重新遍历
			i = det
		}


	}
	//输入：s = "babad"
	//输出："bab"
	//解释："aba" 同样是符合题意的答案。

	// abababa

	// 1234321

	return res
}
