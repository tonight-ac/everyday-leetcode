package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
}

var l, r, mid int
// 二分法就可以
func searchInsert(nums []int, target int) int {
	l, r, mid = 0, len(nums) - 1, 0

	for l <= r {
		mid = (l + r) >> 1
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	return l
}
