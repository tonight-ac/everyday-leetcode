package main

import "sort"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 三个起码选一个，最多选三个，要求和最大
	nums := []int{ root.Val }
	if root.Left != nil { nums = append(nums, maxPathSum(root.Left)) }
	if root.Right != nil { nums = append(nums, maxPathSum(root.Right)) }
	sort.Ints(nums)
	if nums[len(nums)-1] <= 0 {
		return nums[len(nums)-1]
	}

	tot := 0
	for i := len(nums)-1; i >= 0 && nums[i] >= 0; i-- {
		tot += nums[i]
	}

	return tot
}
