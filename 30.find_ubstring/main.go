package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
}

func findSubstring(s string, words []string) []int {
	// 把words组装成一个map，方便搜索
	m := make(map[string]interface{}, len(words))
	for _, v := range words {
		m[v] = nil
	}

	// 长度相同 的单词 words 1 <= words.length <= 5000
	length := len(words[0])

	res := make([]int, 0)
	count := 0
	for i := 0; i <= len(s)-length; {
		if _, ok := m[s[i:i+length]]; ok {
			count++
			i += length
		} else {
			count = 0
			i++
		}

		// 1 1 1 2 2 2 3 3 3
		// 0 1 2 3 4 5 6 7 8
		// 3
		// i-(len(words)-1)*length
		if count == len(words) {
			// 下标计算得到
			res = append(res, i-len(words)*length)
		}
	}

	return res
}
