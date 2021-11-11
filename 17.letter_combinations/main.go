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

// 使用字符串数组，提前创建好映射关系
var m = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// 最多4位，组合出256种情况
var res = make([]string, 256)
// 为了节约内存，开了固定长度res，使用end计算每轮结果最后的下标位置
var end int

// 尝试做一个递归
func letterCombinations(digits string) []string {
	if digits == "" { return nil }

	// 初始化下标，不然全局变量在case之间会互相污染
	end = 0

	recursion(digits, "")

	return res[:end]
}

// TODO 可以改strings.Builder 要在进入之后下一轮函数里做增加
func recursion(digits string, s string) {
	// 递归结束条件
	if digits == "" {
		res[end] = s
		end++
		return
	}

	for _, v := range m[digits[0]] {
		// 插入元素
		s += v
		// 递归
		recursion(digits[1:], s)
		// 删除元素
		s = s[:len(s)-1]
	}
}
