package main

import "fmt"

func main() {
	fmt.Println(numDecodings("226"))
}

// 换一个dp的方式
func numDecodings(s string) int {
	list := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if i>=1 && s[i] >= '1' && s[i] <= '9' { list[i] += list[i-1] + 1 }
		if i>=2 && s[i] != '0' && (s[i-1]-'0')*10+(s[i]-'0') <= 26 { list[i] += list[i-2] + 1 }
		if list[i] == 0 { break }
	}

	return list[len(s)-1]
}

// 递归解决，递归超时了
// case "111111111111111111111111111111111111111111111"
//var res int
//func numDecodings(s string) int {
//	res = 0
//	recursion(s)
//	return res
//}
//
//func recursion(s string) {
//	if len(s) == 0 {
//		res++
//		return
//	}
//
//	if s[0] != '0' {
//		recursion(s[1:])
//	}
//	if len(s) > 1 && s[0] != '0' && (s[0]-'0')*10+(s[1]-'0') <= 26 {
//		recursion(s[2:])
//	}
//}