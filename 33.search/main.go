package main

import "fmt"

//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}

func search(nums []int, target int) int {
	// 先通过二分法判断峰值位置
	mid := 0
	for l, r := 0, len(nums)-1; l <= r; {
		mid = (l + r) / 2
		if nums[mid] > nums[mid+1] {
			break
		} else if nums[mid] < nums[mid+1] {
			l = mid + 1
		}
	}

	fmt.Println(mid)

	// 然后再左侧或右侧再二分搜索
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	return 0
}
