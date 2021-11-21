package main

import "fmt"

func main() {
	fmt.Println(search([]int{1,0,1,1,1}, 0))
}
// TODO 未完成
// 类似33题，只不过允许重复
func search(nums []int, target int) bool {
	// small 必然刚好大于 large
	small, large := nums[0], nums[len(nums)-1]

	// 不满足要求 说明有顺序 直接二分法就可以
	if large > small {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	// 二分法查询反转位置
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= large { // mid在旋转点右侧
			r = mid - 1
		} else if nums[mid] >= small { // mid在旋转点左侧
			l = mid + 1
		}
	}

	fmt.Println(l, r)
	// 4,5,6,7,0,1,2
	// r目前指向7 l目前指向0

	// 在前半段
	if target >= small {
		return binarySearch(nums, 0, r, target)
	}

	// 在后半段
	return binarySearch(nums, l, len(nums) - 1, target)
}

func binarySearch(nums []int, l, r int, target int) bool {
	for l <= r {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
