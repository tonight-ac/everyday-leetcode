package main

import "fmt"

var m = map[int32]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var res = make([]string, 0)

// 尝试做一个递归
func letterCombinations(digits string) []string {
	// 递归结束条件
	if digits == "" {
		return nil
	}

	com := m[int32(digits[0])]

	for _, v := range com {
		// 插入元素
		fmt.Println(v)
		// 递归
		_ = letterCombinations(digits[1:])
		// 删除元素

	}

	return res
}
