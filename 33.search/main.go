package main

import "fmt"

//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
func main() {
	fmt.Println(search([]int{4,5,6,7}, 4))
}

//输入：
//[4,5,6,7,0,1,2]
//3
//输出：
//0
//预期结果：
//-1
func search(nums []int, target int) int {
	// 二分法查询反转位置
	small, large := nums[0], nums[len(nums)-1]

	// 直接二分法就可以
	if large >= small {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] <= large { // mid在旋转点右侧
			r = mid - 1
		} else if nums[mid] >= small { // mid在旋转点左侧
			l = mid + 1
		}
	}

	fmt.Println(l, r)

	return -1
}

func binarySearch(nums []int, l, r int, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
