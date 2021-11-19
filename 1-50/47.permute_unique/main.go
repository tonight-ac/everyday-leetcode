package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
}

var list []int
var res [][]int
// 在46的基础上，做一些微调
func permuteUnique(nums []int) [][]int {
	list = nums

	res = make([][]int, 0)

	recursion([]int{}, make(map[int]bool))

	return res
}

func recursion(r []int, m map[int]bool) {
	if len(r) == len(list) {
		res = append(res, append(make([]int, 0), r...))
		return
	}

	// 在这层栈调用里使用一个局部变量map
	uni := make(map[int]bool)
	for i := 0; i < len(list); i++ {
		if !m[i] && !uni[list[i]] {
			uni[list[i]] = true // 进行去重，不允许这层进行重复的尝试，所以每层都不会出现重复的
			r = append(r, list[i])
			m[i] = true
			recursion(r, m)
			r = r[:len(r)-1]
			m[i] = false
		}
	}
}
