package main

import "fmt"

//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}

func search(nums []int, target int) int {
	// 先通过二分法判断峰值位置
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		fmt.Println(l, r)
		fmt.Println(mid)
		if target < nums[mid] {
			if nums[l] < target {

			}
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	// 然后再左侧或右侧再二分搜索
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		if target < nums[mid] {
			if nums[l] < target {

			}
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	return 0
}
