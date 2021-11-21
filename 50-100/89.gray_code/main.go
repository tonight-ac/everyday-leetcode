package main

import "fmt"

func main() {
	fmt.Println(grayCode(2))
}

// 使用一个递归搜索
// 判断条件是两个数异或能被2整除
var nn int
// 做递归使用1 <= n <= 16
//var list = []int{1,1<<1,1<<2,1<<3,1<<4,1<<5,1<<6,1<<7,1<<8,1<<9,1<<10,1<<11,1<<12,1<<13,1<<14,1<<16}
var uni map[int]bool
var res []int
func grayCode(n int) []int {
	nn = n

	uni = make(map[int]bool)

	res = make([]int, 0)

	uni[0] = true
	recursion(append(make([]int, 0), 0))

	return res
}

func recursion(list []int) bool {
	if len(list) == 1<<nn {
		if (list[len(list)-1] & 0) % 2 == 0 {
			res = append(res, list...)
			return true
		}
		return false
	}

	// 上一个数字是多少，直接决定当前数字是多少
	cur := list[len(list)-1]
	ok := false
	for i := 0; i < nn; i++ {
		// 一位一位遍历
		temp := 0
		if pos := 1<<i; cur & pos == 0 { // 说明cur这位是0
			temp = cur + pos
		} else { // 说明cur这位是1
			temp = cur - pos
		}
		// 不允许重复
		if !uni[temp] {
			uni[temp] = true
			if ok = recursion(append(list, temp)); ok {
				break
			}
			uni[temp] = false
		}
	}

	return ok
}

