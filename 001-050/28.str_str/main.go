package main

import "fmt"

func main() {
	strStr("", "AAAAAAAA")
	// ABCABCD
	// [0 0 0 1 2 3 0]
}

//输入：haystack = "hello", needle = "ll"
//输出：2
// 简单来说就是最经典的kmp算法，课本上都会讲
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)

	if m == 0 { return 0 }

	//输入：haystack = "hello", needle = "ll"
	//输出：2

	// 计算next数组
	// next表示的含义是，如果在needle在这个位置失去和haystack的匹配了，如果保持haystack的指针不变，移动needle指针的位置
	// ABCABCD
	next := make([]int, m)
	// 自己先和自己匹配一遍，求出next数组
	// ABCABCABCD
	// [0 0 0 1 2 3 4 5 6 0]
	// AAAAAAAA
	// [0 1 2 3 4 5 6 7]
	for i, j := 1, 0; i < m; i++ {
		fmt.Println(i, j)
		for j > 0 && needle[i] != needle[j] {
			j = next[j - 1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)

	for i, j := 0, 0; i < n; i++ {
		// 不是归0的话，就一直循环回退，直到找到合适的haystack[i] != needle[j]
		for j > 0 && haystack[i] != needle[j] {
			j = next[j - 1] // j失败，选取j-1的位置
		}
		// 相同，needle向后移动匹配下一个
		if haystack[i] == needle[j] { j++ }
		// 匹配上了，找到了，返回
		if j == m { return i - m + 1 }
	}

	return -1
}

//func strStr(haystack string, needle string) int {
//	if needle == "" {
//		return 0
//	}
//
//	m := make(map[string]int)
//
//	for i := 0; i <= len(haystack) - len(needle); i++ {
//		if _, ok := m[haystack[i:i+len(needle)]]; !ok {
//			m[haystack[i:i+len(needle)]] = i
//		}
//	}
//
//	if idx, ok := m[needle]; ok {
//		return idx
//	}
//
//	return -1
//}
