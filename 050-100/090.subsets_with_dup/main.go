package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{4,4,4,1,4}))
}

// 类似78题，但存在递进
// 和78题的递进关系类似46题和47题
var res [][]int
var n []int
var uni map[int]bool
func subsetsWithDup(nums []int) [][]int {
	n = nums

	res = make([][]int, 0)

	uni = make(map[int]bool)

	// 先排个顺序
	sort.Ints(nums)

	recursion(make([]int, 0), 0)

	return res
}

func recursion(s []int, idx int) {
	if idx == len(n) {
		res = append(res, append(make([]int, 0), s...))
		return
	}

	// 如果已经存在了，则必须选择当前元素
	// 例如[0]2,[1]2,[2]2
	// 因为上一个存在的话，必然有不存在的case
	if !uni[n[idx]] {
		// 不选择
		recursion(s, idx+1)
	}

	uni[n[idx]] = true
	recursion(append(s, n[idx]), idx+1)
	uni[n[idx]] = false
}
