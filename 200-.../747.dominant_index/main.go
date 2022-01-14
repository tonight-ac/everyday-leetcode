package main

import "fmt"

//给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。
//
//请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。
//
//示例 1：
//
//输入：nums = [3,6,1,0]
//输出：1
//解释：6 是最大的整数，对于数组中的其他整数，6 大于数组中其他元素的两倍。6 的下标是 1 ，所以返回 1 。

//1 <= nums.length <= 50
//0 <= nums[i] <= 100
//nums 中的最大元素是唯一的

func main() {
	fmt.Println(dominantIndex([]int{1}))
}

func dominantIndex(nums []int) int {
	// 寻找到第一大和第二大
	tar, first, second := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > first {
			second = first
			first = nums[i]
			tar = i
		} else if nums[i] > second {
			second = nums[i]
		}
	}

	if first >= 2 * second {
		return tar
	}

	return -1
}
