package main

import (
	"fmt"
	"strings"
)

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

//func isScramble(s1, s2 string) bool {
//	// 对m数组初始化 长度 n n n+1 数值 -1
//	n := len(s1)
//	m := make([][][]int8, n)
//	for i := range m {
//		m[i] = make([][]int8, n)
//		for j := range m[i] {
//			m[i][j] = make([]int8, n+1)
//			for k := range m[i][j] {
//				m[i][j][k] = -1
//			}
//		}
//	}
//
//	// 第一个字符串从 i1 开始，第二个字符串从 i2 开始，子串的长度为 length
//	// 和谐返回 1，不和谐返回 0
//	var dfs func(i1, i2, length int) int8
//	dfs = func(i1, i2, length int) (res int8) {
//		d := &m[i1][i2][length]
//		if *d != -1 {
//			return *d
//		}
//		defer func() { *d = res }()
//
//		// 判断两个子串是否相等
//		x, y := s1[i1:i1+length], s2[i2:i2+length]
//		if x == y {
//			return 1
//		}
//
//		// 判断是否存在字符 c 在两个子串中出现的次数不同
//		freq := [26]int{}
//		for i, ch := range x {
//			freq[ch-'a']++
//			freq[y[i]-'a']--
//		}
//		for _, f := range freq[:] {
//			if f != 0 {
//				return 0
//			}
//		}
//
//		// 枚举分割位置
//		for i := 1; i < length; i++ {
//			// 不交换的情况
//			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
//				return 1
//			}
//			// 交换的情况
//			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
//				return 1
//			}
//		}
//
//		return 0
//	}
//	return dfs(0, 0, n) == 1
//}

//"acddaaaadbcbdcdaccabdbadccaaa"
//"adcbacccabbaaddadcdaabddccaaa"
func main() {
	//fmt.Println(isScramble("abcd", "badc"))
	//fmt.Println(isScramble("abcde", "caebd"))
	//fmt.Println(isScramble("a", "a"))
	fmt.Println(isScramble("acddaaaadbcbdcdaccabdbadccaaa", "adcbacccabbaaddadcdaabddccaaa"))
}

// 又超时了nm的
var m1, m2 [][]string
func isScramble(s1 string, s2 string) bool {
	//[
	//[10000000000000000000000000 11000000000000000000000000 11100000000000000000000000 11110000000000000000000000]
	//[ 01000000000000000000000000 01100000000000000000000000 01110000000000000000000000]
	//[  00100000000000000000000000 00110000000000000000000000]
	//[   00010000000000000000000000]
	//]

	m1 = make([][]string, len(s1))
	for i := 0; i < len(s1); i++ {
		m1[i] = make([]string, len(s1))
		for j := i; j < len(s1); j++ {
			m1[i][j] = hash(s1[i:j+1])
		}
	}

	//[
	//[01000000000000000000000000 11000000000000000000000000 11010000000000000000000000 11110000000000000000000000]
	//[ 10000000000000000000000000 10010000000000000000000000 10110000000000000000000000]
	//[  00010000000000000000000000 00110000000000000000000000]
	//[   00100000000000000000000000]
	//]

	m2 = make([][]string, len(s2))
	for i := 0; i < len(s2); i++ {
		m2[i] = make([]string, len(s2))
		for j := i; j < len(s2); j++ {
			m2[i][j] = hash(s2[i:j+1])
		}
	}

	return recursion(0, len(s1)-1, 0, len(s2)-1)
}
// 0 3 0 3
// 0 1 2
// 0 1 0 1
// 0
// 搞一个记忆化搜索
// 把结果存储起来
// 首先j1-i1==j2-i2
func recursion(i1, j1, i2, j2 int) bool {
	if i1 == j1 && i2 == j2 {
		return m1[i1][j1] == m2[i2][j2]
	}

	for i := 0; i < j1-i1; i++ {
		// 不交换
		if m1[i1][i1+i] == m2[i2][i2+i] && m1[i1+i+1][j1] == m2[i2+i+1][j2] {
			if ok := recursion(i1, i1+i, i2, i2+i) && recursion(i1+i+1, j1, i2+i+1, j2); ok {
				return true
			}
		}
		// 交换
		if m1[i1][i1+i] == m2[j2-i][j2] && m1[i1+i+1][j1] == m2[i2][j2-i-1] {
			if ok := recursion(i1, i1+i, j2-i, j2) && recursion(i1+i+1, j1, i2, j2-i-1); ok {
				return true
			}
		}
	}

	return false
}
// hash函数，计算26个字母的出现次数
// 因为1 <= s1.length <= 30
// 所以最多30个相同的字母，一个ascii足够存储
func hash(s string) string {
	list := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		list[s[i]-'a']++
	}

	var b strings.Builder
	for _, v := range list {
		b.WriteByte(v+'0')
	}

	return b.String()
}
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

//"abcdbdacbdac"
//"bdacabcdbdac"
//func isScramble(s1 string, s2 string) bool {
//	stack := make([]byte, 0)
//	i, j := 0, 0
//	for {
//		if i < len(s1) && j < len(s2) && s1[i] == s2[j] {
//			for i < len(s1) && j < len(s2) && s1[i] == s2[j] {
//				i++
//				j++
//			}
//		} else if j < len(s2) && len(stack) != 0 {
//			for j < len(s2) && len(stack) != 0 {
//				if stack[len(stack)-1] == s2[j] {
//					// 尾部
//					stack = stack[:len(stack)-1]
//					j++
//				} else if stack[0] == s2[j] {
//					// 头部
//					stack = stack[1:]
//					j++
//				} else {
//					return false
//				}
//			}
//		} else if i < len(s1) {
//			stack = append(stack, s1[i])
//			i++
//		} else {
//			break
//		}
//	}
//	return len(stack) == 0 && j == len(s2)
//}

// 用一个栈解决，借鉴了84的思路
// 不是所有题都需要用m，很多时候向前搜索都需要用到栈
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

