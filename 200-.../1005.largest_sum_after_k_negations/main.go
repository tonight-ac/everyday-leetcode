package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{2,-3,-1,5,-4}, 2))
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	if nums[0] < 0 {
		// 处理存在负数的
		for i := 0; i < len(nums) && nums[i] < 0 && k > 0; i, k = i + 1, k - 1 {
			nums[i] = -nums[i]
		}
		// 还有多余的k，那么现在必然都是正数，按照全部大等于0的情况处理
		if k > 0 {
			gteZero(nums, k, true)
		}
	} else {
		gteZero(nums, k, false)
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res
}

func gteZero(nums []int, k int, needSort bool) {
	if needSort { sort.Ints(nums) }
	if nums[0] == 0 {
		// 有0直接对0做k次变形
	} else if nums[0] > 0 {
		// 对第一个元素做k次变形
		if k % 2 == 1 {
			nums[0] = -nums[0]
		}
	}
}