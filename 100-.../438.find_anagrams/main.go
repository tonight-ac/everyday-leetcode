package main

import "fmt"

//"abacbabc"
//"abc"
func main() {
	fmt.Println(findAnagrams("abacbabc", "abc"))
}

// 借鉴49题
// 搞一个滑动窗口
var uni [26]int16
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) { return nil }

	for i := 0; i < len(uni); i++ {
		uni[i] = 0
	}
	for i := 0; i < len(p); i++ {
		uni[p[i]-'a']++
	}

	res := make([]int, 0)
	// 0 1 2 3
	count := 0
	for i := 0; i < len(s); i++ {
		for uni[s[i]-'a'] < 1 && count > 0 {
			// 归还
			uni[s[i-count]-'a']++
			count--
		}
		if uni[s[i]-'a'] > 0 {
			uni[s[i]-'a']--
			count++
		}

		if count == len(p) {
			// 记录下标
			res = append(res, i+1-count)
			// 将首个清除掉
			uni[s[i+1-count]-'a']++
			count--
		}
	}

	return res
}

