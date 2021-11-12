package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// 2 <= n <= 8
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]

var tot int
var res []string
func generateParenthesis(n int) []string {
	res = make([]string, 0)
	tot = n
	recursion(0, 0, "")
	return res
}

func recursion(left, right int, s string) {
	if len(s) == 2 * tot {
		res = append(res, s)
		return
	}

	if left < tot {
		s += "("
		recursion(left+1, right, s)
		s = s[:len(s)-1]
	}
	if left <= tot && right < left {
		s += ")"
		recursion(left, right+1, s)
		s = s[:len(s)-1]
	}
}
