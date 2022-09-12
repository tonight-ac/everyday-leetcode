package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{2,2,2,2,2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	// 对数组进行排序
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 去重 因为i所遍历的后续内容，必定是i-1所遍历的后续内容的子集
		if i > 0 && nums[i] == nums[i-1] { continue }

		res = append(res, threeSum(nums[i], nums[i+1:], target)...)
	}

	return res
}

func threeSum(num int, nums []int, target int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		// 去重 因为i所遍历的后续内容，必定是i-1所遍历的后续内容的子集
		if i > 0 && nums[i] == nums[i-1] { continue }

		// 锁定一个，找到包含它的所有tuple
		for l, r := i+1, len(nums)-1; l < r; {
			sum := num+nums[i]+nums[l]+nums[r]
			if sum == target { // 已经match
				res = append(res, []int{num, nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] { l++ } // 目前已经match了，再计算也是重复的，移动到最后一个重复元素
				for l < r && nums[r] == nums[r-1] { r-- } // 目前已经match了，再计算也是重复的，移动到最后一个重复元素
				l++ // 此时再次循环还是重复的，还可以发生match，所以再移动一个
				r-- // 此时再次循环还是重复的，还可以发生match，所以再移动一个
				// 如果都未发生重复，也本身就需要移动，而且由于有序性，不可能同样的数字还有多解法，两个都移动
			} else if sum < target { // 三数之和为负，说明较小，需要向大的方向移动
				l++
			} else { // 三数之和为正，说明较大，需要向小的方向移动
				r--
			}
		}
	}

	return res
}
