package main

import (
	"fmt"
)

// #a#a#b#
// #b#a#a#
// a a b aa




func main() {
	//fmt.Println(countSubstrings("efe"))
	fmt.Println(partition("efe"))
}

// 先找到所有回文子串的下标，然后搞搜索
var m map[int]map[int]interface{}
var res [][]string
var ss string
func partition(s string) [][]string {
	ss = s
	m = countSubstrings(s)
	res = make([][]string, 0)

	recursion(0, make([]string, 0))

	return res
}

func recursion(i int, list []string) {
	// 终止条件
	if i == len(ss) {
		res = append(res, append(make([]string, 0), list...))
		return
	}

	for k, _ := range m[i] {
		recursion(k, append(list, ss[i:k]))
	}
}

// 找到所有回文子串的下标
// 需要参考647题
func countSubstrings(s string) map[int]map[int]interface{} {
	mm := make(map[int]map[int]interface{})
	for i := 0; i < len(s); i++ { mm[i] = make(map[int]interface{}) }

	// 单中心
	for i := 0; i < len(s); i++ {
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			mm[l][r+1] = nil
		}
	}

	// 双中心
	for i := 1; i < len(s); i++ {
		for l, r := i-1, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			mm[l][r+1] = nil
		}
	}

	return mm
}
