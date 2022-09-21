package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// 2 <= n <= 8
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]

// 保存公共结果，减少本题内存消耗
var tot int
var res []string

// 做一个递归操作
func generateParenthesis(n int) []string {
	res = make([]string, 0)
	tot = n
	recursion(0, 0, "")
	return res
}

// 递归函数
func recursion(left, right int, s string) {
	// 终止条件，字符串s长度符合要求
	if len(s) == 2 * tot {
		res = append(res, s)
		return
	}

	// 如果左号已经出现次数小于总长度一半，可以添加
	if left < tot {
		s += "("
		recursion(left+1, right, s)
		s = s[:len(s)-1]
	}
	// 如果左号已经出现次数小于等于总长度一半，且右号数量小于左号数量
	if right < left {
		s += ")"
		recursion(left, right+1, s)
		s = s[:len(s)-1]
	}
}
