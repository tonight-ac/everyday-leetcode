package main

import "fmt"

// 需要自底向上搜索
//var tar string
//var res bool
//func isScramble(s1 string, s2 string) bool {
//	if s1 == s2 { return true }
//
//	tar = s2
//
//	res = false
//
//	return res
//}
//
//func recursion(s string, ss []string) {
//	if len(ss) == len(tar) {
//		if strings.Join(ss, "") == tar {
//			res = true
//		}
//		return
//	}
//
//
//}

// 下面的case超时了
//"ccabcbabcbabbbbcbb"
//"bbbbabccccbbbabcba"
// 自顶向下搜索是不可以的

func main() {
	fmt.Println(isScramble("abc", "cba"))
}

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 { return true }

	// 0 1 2 3 4 5
	for i := 1; i < len(s1); i++ {
		// 不交换
		if ok := isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]); ok {
			return true
		}
		// 交换
		if ok := isScramble(s1[i:], s2[:len(s2)-i]) && isScramble(s1[:i], s2[len(s2)-i:]); ok {
			return true
		}
	}

	return false
}
