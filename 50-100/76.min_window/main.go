package main

import "fmt"

//输入：s = "ADOBECODEBANC", t = "ABC"
// ABCBAC
//输出："BANC"
// TODO
func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须*不少于* t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。

func minWindow(s string, t string) string {
	// 将t整理成map，方便处理
	m := make(map[byte]int)
	for _, v := range t { m[byte(v)]++ }

	ch := make([]byte, 0)
	pos := make([]int, 0)
	for i, v := range s {
		// 找出所有符合要求的字符，以及他们在s中的下标
		if m[byte(v)] > 0 {
			ch = append(ch, byte(v))
			pos = append(pos, i)
		}
	}

	// 找到最小子串的启始下标，和最小长度
	start, idx, min, count := 0, 0, 0, 0
	for i := 0; i < len(ch); i++ {
		// 如果大于0，那选择ch[i]没错，贪心算法
		if m[ch[i]] > 0 {
			count++
		}
		m[ch[i]]--

		// 到此匹配上了
		if count == len(t) {
			// 记录最小值
			if min > pos[i]-pos[start] {
				min = pos[i]-pos[start]
				idx = start
			}
			// 删除第start个
			m[ch[start]]++
			// 归还count
			if m[ch[start]] > 0 { count-- }
			// 移动start到下一个
			start++

		}
	}

	return s[pos[idx]:pos[idx]+min]
}
