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
	//fmt.Println(isScramble("abcd", "badc"))
	//fmt.Println(isScramble("abcde", "caebd"))
	//fmt.Println(isScramble("a", "a"))
	fmt.Println(isScramble("abcdbdacbdac", "bdacabcdbdac"))
}

//"abcdbdacbdac"
//"bdacabcdbdac"
func isScramble(s1 string, s2 string) bool {
	stack := make([]byte, 0)
	i, j := 0, 0
	for {
		if i < len(s1) && j < len(s2) && s1[i] == s2[j] {
			for i < len(s1) && j < len(s2) && s1[i] == s2[j] {
				i++
				j++
			}
		} else if j < len(s2) && len(stack) != 0 {
			for j < len(s2) && len(stack) != 0 {
				if stack[len(stack)-1] == s2[j] {
					// 尾部
					stack = stack[:len(stack)-1]
					j++
				} else if stack[0] == s2[j] {
					// 头部
					stack = stack[1:]
					j++
				} else {
					return false
				}
			}
		} else if i < len(s1) {
			stack = append(stack, s1[i])
			i++
		} else {
			break
		}
	}
	return len(stack) == 0 && j == len(s2)
}
// 用一个栈解决，借鉴了84的思路
// 不是所有题都需要用dp，很多时候向前搜索都需要用到栈
//"abcde"
//"caebd"

// abcd
// dcba
//func isScramble(s1 string, s2 string) bool {
//	stack := make([]byte, 0)
//	//s1, s2 = s2, s1
//	i, j := 0, 0
//	for i < len(s1) && j < len(s2) {
//		if s1[i] == s2[j] {
//			i++
//			j++
//			continue
//		} else {
//			if len(stack) == 0 {
//				stack = append(stack, s1[i])
//				i++
//				continue
//			} else {
//				for len(stack) != 0 {
//					if stack[len(stack)-1] == s2[j] {
//						// 尾部
//						stack = stack[:len(stack)-1]
//						j++
//						continue
//					} else if stack[0] == s2[j] {
//						// 头部
//						stack = stack[1:]
//						j++
//						continue
//					} else {
//						return false
//					}
//				}
//			}
//		}
//	}
//
//	for len(stack) != 0 && j < len(s2) {
//		if stack[len(stack)-1] == s2[j] {
//			stack = stack[:len(stack)-1]
//			j++
//		} else if stack[0] == s2[j] {
//			stack = stack[1:]
//			j++
//		} else {
//			return false
//		}
//	}
//
//	return len(stack) == 0 && j == len(s2)
//}

//func isScramble(s1 string, s2 string) bool {
//	if s1 == s2 { return true }
//
//	// 0 1 2 3 4 5
//	for i := 1; i < len(s1); i++ {
//		// 不交换
//		if ok := isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]); ok {
//			return true
//		}
//		// 交换
//		if ok := isScramble(s1[i:], s2[:len(s2)-i]) && isScramble(s1[:i], s2[len(s2)-i:]); ok {
//			return true
//		}
//	}
//
//	return false
//}
