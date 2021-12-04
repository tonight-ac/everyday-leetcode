package main

import (
	"fmt"
	"strings"
)

// #a#a#b#
// #b#a#a#
// a a b aa




func main() {
	fmt.Println(findAllSubPalindrome("efe"))
	//fmt.Println(partition("aab"))
}

// 先找到所有回文子串的下标，然后搞搜索
var m map[int]map[int]interface{}
var res [][]string
var ss string
func partition(s string) [][]string {
	ss = s
	m = findAllSubPalindrome(s)
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

// 找到所有回文子串的下标 TODO
func findAllSubPalindrome(s string) map[int]map[int]interface{} {
	res := make(map[int]map[int]interface{})

	ss := "#" + strings.Join(strings.Split(s, ""), "#") + "#"

	i, j := 0, 0

	for i = 0; i < len(ss); i++ {
		for j = i + 1; j < len(ss); j++ {
			if 2 * i - j < 0 || ss[j] != ss[2 * i - j] {
				break
			}
		}

		realI, realJ := (2*i-(j-1))/2, (j-1)/2
		if res[realI] == nil {
			res[realI] = make(map[int]interface{})
		}
		if realI != realJ {
			res[realI][realJ] = nil
		}
	}

	return res
}


//func isPalindrome(s string) bool {
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		if s[i] != s[j] { return false }
//	}
//
//	return true
//}