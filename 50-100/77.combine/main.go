package main

import "fmt"

func main() {
	fmt.Println(combine(1,1))
}

// TODO 好慢的方法，看有没有快一点的
var nn, kk int
var res [][]int
var uni [21]bool // 1<= n <= 20 开21空间即可
func combine(n int, k int) [][]int {
	nn, kk = n, k

	uni = [21]bool{}

	res = make([][]int, 0)

	recursion(make([]int, 0), 1)

	return res
}

func recursion(com []int, idx int) {
	if len(com) == kk {
		res = append(res, append(make([]int, 0), com...))
		return
	}

	if idx > nn {
		return
	}

	for i := idx; i <= nn; i++ {
		if !uni[i] && (len(com) == 0 || i > com[len(com)-1]) {
			uni[i] = true
			recursion(append(com, i), idx+1)
			uni[i] = false
		}
	}
}