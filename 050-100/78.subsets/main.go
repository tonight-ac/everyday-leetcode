package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}

// 内存消耗较大
var res [][]int
var n []int
func subsets(nums []int) [][]int {
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

	// 不选择
	recursion(s, idx+1)
	// 选择
	recursion(append(s, n[idx]), idx+1)
}
