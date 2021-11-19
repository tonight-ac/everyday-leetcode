package main

import "fmt"

func main() {
	fmt.Println(getPermutation(4, 9))
}

var nn, kk int
var res string
var uni [10]bool // 1<= n <= 9 开10空间即可
func getPermutation(n int, k int) string {
	nn, kk = n, k

	uni = [10]bool{}

	recursion("")

	return res
}

func recursion(s string) {
	if len(s) == nn {
		kk-- // 完成一次做一次扣减
		if kk == 0 { // 这就是第k次，记录下答案
			res = s
		}
		return
	}
	for i := 1; i <= nn; i++ {
		// uni去重用的
		if !uni[i] && kk > 0 { // kk > 0 减枝操作，已经拿到答案了，不需要再尝试了
			uni[i] = true
			recursion(s+string(rune(i+'0')))
			uni[i] = false
		}
	}
}