package main

import "fmt"

//"barfoofoobarthefoobarman"
//["bar","foo","the"]
//输出：
//[9]
//预期结果：
//[6,9,12]

// 应该存在类似kmp算法的类似解法
func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"word","good","best","good"}))
}
//"wordgoodgoodgoodbestword"
//["word","good","best","good"]
//输出：
//[]
//预期结果：
//[8]
// 递归解法
var m map[string]int
var length int
func findSubstring(s string, words []string) []int {
	// 几步走？
	// 1 当前是否匹配

	// words是可能重复的
	m = make(map[string]int, len(words))
	for _, v := range words {
		m[v] = m[v] + 1
	}

	length = len(words[0])

	size := len(words)*length

	res := make([]int, 0)

	for i := 0; i <= len(s)-size; i++ {
		// 处理i~i+size的字符
		// 长度不然是这么长，我们不判断什么符合，我们只判断什么是异常情况
		if recursion(s[i:i+size]) {
			res = append(res, i)
		}
	}

	return res
}

func recursion(s string) bool {
	if s == "" { return true }

	ss := s[:length]
	if m[ss] - 1 < 0 {
		return false
	}

	m[ss]--
	res := recursion(s[length:])
	m[ss]++

	return res
}

//"barfoofoobarthefoobarman"
//["bar","foo","the"]
//输出：
//[9,12]
//预期结果：
//[6,9,12]
//func main() {
//	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
//}
//
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
//		//fmt.Println(i)
//		//fmt.Println(s[i:i+length])
//		if uni[s[i:i+length]] {
//			count = 0
//			uni = make(map[string]bool, len(words))
//		}
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
