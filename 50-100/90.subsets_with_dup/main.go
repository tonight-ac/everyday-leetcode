package main

import "fmt"

func main() {
	fmt.Println(subsetsWithDup([]int{1,2,2}))
}
// TODO 未完成
// 类似78题，但存在递进
// 和78题的递进关系类似46题和47题
var res [][]int
var n []int
func subsetsWithDup(nums []int) [][]int {
	n = nums

	res = make([][]int, 0)

	recursion(make([]int, 0), 0)

	return res
}

func recursion(s []int, idx int) {
	if idx == len(n) {
		res = append(res, append(make([]int, 0), s...))
		return
	}

	// 当前等于之前的，直接忽略掉
	if idx == 0 || n[idx] != n[idx-1] {
		// 选择
		recursion(append(s, n[idx]), idx+1)
	} else if len(s) > 0 && n[idx] == s[len(s)-1] {
		// 选择
		recursion(append(s, n[idx]), idx+1)
	}

	// 不选择
	recursion(s, idx+1)
}
