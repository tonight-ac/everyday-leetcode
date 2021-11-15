package main

import "fmt"

//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
func main() {
	fmt.Println(search([]int{1, 3}, 0))
}

//输入：
//[4,5,6,7,0,1,2]
//3
//输出：
//0
//预期结果：
//-1
func search(nums []int, target int) int {
	if len(nums) == 0 { return -1 }

	//gapL, gapR := nums[0], nums[len(nums)-1]
	//
	//for l, r := 0, len(nums)-1; l <= r; {
	//	mid := (l + r) / 2
	//	if target < nums[mid] {
	//
	//		r = mid - 1
	//	} else if target > nums[mid] {
	//		l = mid + 1
	//	} else {
	//		return mid
	//	}
	//}

	// 先通过二分法判断峰值位置
	//l, r, mid := 0, 0, 0
	//for l, r = 0, len(nums)-1; l <= r; {
	//	mid = (l + r) / 2
	//	// 这种情况说明，该序列为生序
	//	if mid + 1 >= len(nums) { break }
	//	if nums[mid] > nums[mid+1] {
	//		break
	//	} else if nums[mid] < nums[mid+1] {
	//		l = mid + 1
	//	}
	//}
	//
	//if mid + 1 >= len(nums) {
	//	l, r = 0, len(nums) - 1
	//}else if target < nums[len(nums)-1] {
	//	l, r = mid + 1, len(nums)-1
	//} else if target > nums[len(nums)-1] {
	//	l, r = 0, mid - 1
	//} else {
	//	return mid
	//}
	//
	//// 然后再左侧或右侧再二分搜索
	//for l <= r {
	//	mid = (l + r) / 2
	//	if target < nums[mid] {
	//		r = mid - 1
	//	} else if target > nums[mid] {
	//		l = mid + 1
	//	} else {
	//		return mid
	//	}
	//}

	return -1
}
