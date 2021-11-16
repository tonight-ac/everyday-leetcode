package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}

var res [][]int
var can []int
func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)

	// 排序避免重复
	sort.Ints(candidates)

	can = candidates

	recursion(len(candidates)-1, target, make([]int, 0))

	return res
}

func recursion(idx int, tar int, list []int) {
	// 查找到一种情况，返回
	if tar == 0 {
		res = append(res, append(make([]int, 0), list...))
		return
	}
	// 越界（无结果）或不是正好相等
	if idx < 0 || tar < 0 {
		return
	}

	// 使用can[idx]
	list = append(list, can[idx])
	recursion(idx, tar - can[idx], list)
	list = list[:len(list)-1]

	// 不使用can[idx]
	recursion(idx - 1, tar, list)
}
