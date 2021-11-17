package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。

//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。

// 类似第10题的解法，不过递归超时了
func isMatch(s string, p string) bool {
	// 如果p已经匹配完了，那就判断一下s有没有匹配完
	if p == "" {
		return s == ""
	}

	if p[0] == '*' {
		// p为'*' 那分为匹配当前s和不匹配当前s两种结果
		return (len(s) > 0 && isMatch(s[1:], p)) || isMatch(s, p[1:])
	}

	// 相等 或者 p为'?'可以匹配任何单个字符
	return len(s) > 0 && (s[0] == p[0] || p[0] == '?') && isMatch(s[1:], p[1:])
}
