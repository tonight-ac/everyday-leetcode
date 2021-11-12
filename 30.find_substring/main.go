package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
}

//"barfoofoobarthefoobarman"
//["bar","foo","the"]
//输出：
//[9]
//预期结果：
//[6,9,12]
// TODO 试一下递归吧

func findSubstring(s string, words []string) []int {
	// 几步走？
	// 1 当前下标处

	return nil
}

func recursion(m map[string]bool, s string, count int) {
	if count == 3 {
		return
	}

	// 当前匹配
	if v, ok := m[s[0:3]]; ok && !v {
		// match 上之后继续向下
		m[s[0:3]] = true
		recursion(m, s[3:], count+1)
		m[s[0:3]] = false
	}

	// 不匹配继续向下
	recursion(m, s[1:], count)

}

//func findSubstring(s string, words []string) []int {
//	// 把words组装成一个map，方便搜索
//	m := make(map[string]bool, len(words))
//	for _, v := range words {
//		m[v] = true
//	}
//
//	// 长度相同 的单词 words 1 <= words.length <= 5000
//	length := len(words[0])
//
//	res := make([]int, 0)
//	count := 0
//	// 循环里新开一个局部去重map
//	uni := make(map[string]bool, len(words))
//	for i := 0; i <= len(s)-length; {
//		if m[s[i:i+length]] && !uni[s[i:i+length]] {
//			uni[s[i:i+length]] = true
//			count++
//			i += length
//		} else {
//			uni = make(map[string]bool, len(words))
//			count = 0
//			i++
//		}
//
//		// 1 1 1 2 2 2 3 3 3
//		// 0 1 2 3 4 5 6 7 8
//		// 3
//		// i-len(words)*length
//		if count == len(words) {
//			// 下标计算得到
//			res = append(res, i-len(words)*length)
//			i = i-len(words)*length+1
//		}
//	}
//
//	return res
//}
