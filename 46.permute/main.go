package main

//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

var list []int
var res [][]int
func permute(nums []int) [][]int {
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

	for i := 0; i < len(list); i++ {
		if !m[list[i]] {
			r = append(r, list[i])
			m[list[i]] = true
			recursion(r, m)
			r = r[:len(r)-1]
			m[list[i]] = false
		}
	}
}
