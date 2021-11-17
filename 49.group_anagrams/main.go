package main

import "strings"

// 需要设计一个hash函数，让字符相同的hash出来的值都一样
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]string)
	for _, v := range strs {
		if m[hash(v)] == nil {
			m[hash(v)] = append(make([]string, 0), v)
		} else {
			m[hash(v)] = append(m[hash(v)], v)
		}
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// hash函数，计算26个字母的出现次数
// 因为0 <= strs[i].length <= 100
// 所以最多100个相同的字母，一个ascii足够存储
func hash(s string) string {
	list := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		list[s[i]-'a']++
	}

	var b strings.Builder
	for _, v := range list {
		b.WriteByte(v)
	}

	return b.String()
}