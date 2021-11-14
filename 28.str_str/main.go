package main

//输入：haystack = "hello", needle = "ll"
//输出：2
// 简单来说就是最经典的kmp算法，课本上都会讲
func strStr(haystack string, needle string) int {

	return 0
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
