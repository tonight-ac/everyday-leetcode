package main

import (
	"fmt"
)

//"2"
//输出：
//["ad","ae","af","bd","be","bf","cd","ce","cf","a","b","c"]
//预期结果：
//["a","b","c"]

func main() {
	fmt.Println(letterCombinations("2"))
}
// TODO 未发题解
var m = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
//var s string
//func letterCombinations(digits string) []string {
//	// 递归结束条件
//	if digits == "" {
//		if s != "" {
//			return []string{s}
//		}
//	}
//
//	res := make([]string, 0)
//
//	for _, v := range m[digits[0]] {
//		// 插入元素
//		s += string(v)
//		// 递归
//		res = append(res, letterCombinations(digits[1:])...)
//		// 删除元素
//		s = s[:len(s)-1]
//	}
//
//	return res
//}
var res []string

// 尝试做一个递归
func letterCombinations(digits string) []string {
	if digits == "" { return nil }
	// 4 4 4 4
	//
	res = make([]string, 0)

	recursion(digits, "")

	return res
}


func recursion(digits string, s string) {
	// 递归结束条件
	if digits == "" {
		res = append(res, s)
		return
	}

	for _, v := range m[digits[0]] {
		// 插入元素
		s += string(v)
		// 递归
		recursion(digits[1:], s)
		// 删除元素
		s = s[:len(s)-1]
	}
}
