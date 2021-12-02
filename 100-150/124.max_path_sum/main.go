package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(maxPathSum(&TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}
//[1,-2,-3,1,3,-2,null,-1]
//      1
//   -2   -3
//  1  3 -2
// -1
// TODO未完成
func maxPathSum(root *TreeNode) int {
	// 三个起码选一个，最多选三个，要求和最大
	var left, right int
	nums := []int{ root.Val }
	if root.Left != nil {
		left = maxPathSum(root.Left)
		nums = append(nums, left)
		nums = append(nums, left + root.Val)
	}
	if root.Right != nil {
		right = maxPathSum(root.Right)
		nums = append(nums, right)
		nums = append(nums, right + root.Val)
	}
	if root.Left != nil && root.Right != nil {
		nums = append(nums, left + right + root.Val)
	}

	sort.Ints(nums)
	if root.Right != nil && root.Right.Val == 3 { fmt.Println(nums) }

	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}
