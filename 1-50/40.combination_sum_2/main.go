package main

import (
	"fmt"
	"sort"
)

//[2,5,2,1,2]
//5

func main() {
	fmt.Println(combinationSum2([]int{2,2,2}, 2))
}

var res [][]int
var can []int
func combinationSum2(candidates []int, target int) [][]int {
	res = make([][]int, 0)

	// 使用map
	m := make(map[int]int)
	for _, v := range candidates {
		m[v] = m[v] + 1
	}

	// 排序避免重复
	sort.Ints(candidates)

	can = candidates

	recursion(len(candidates)-1, target, make([]int, 0), m)

	return res
}

func recursion(idx int, tar int, list []int, m map[int]int) {
	// 查找到一种情况，返回
	if tar == 0 {
		res = append(res, append(make([]int, 0), list...))
		return
	}

	// 越界（无结果）或不是正好相等
	if idx < 0 || tar < 0 {
		return
	}

	// 如果尝试了m[can[idx]]这条路，并且can[idx] != can[idx-1]，就没必要试了
	// 如果没有尝试
	// 使用can[idx]

	if m[can[idx]] > 0 {
		m[can[idx]]--
		list = append(list, can[idx])
		recursion(idx - 1, tar - can[idx], list, m)
		m[can[idx]]++
		list = list[:len(list)-1]
	}

	// 此处要用循环，如果出现重复，你不能在上面的if中不进入，转而选择下一个完全相同的再次进行尝试
	// 这样势必会导致重复
	for idx > 0 && can[idx] == can[idx - 1] {
		idx--
	}
	recursion(idx - 1, tar, list, m)
}
