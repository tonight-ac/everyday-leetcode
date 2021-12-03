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
		Val: -3,
	}))
	//fmt.Println(maxPathSum(&TreeNode{
	//	Val: -10,
	//	Left: &TreeNode{
	//		Val: 9,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 15,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//}))
}
//[1,-2,-3,1,3,-2,null,-1]
//      1
//   -2   -3
//  1  3 -2
// -1
// TODO未完成
var res int
func maxPathSum(root *TreeNode) int {
	res = root.Val

	recursion(root)

	return res
}

func recursion(root *TreeNode) int {
	// 三个起码选一个，最多选三个，要求和最大
	var left, right int
	nums := []int{ root.Val }
	if root.Left != nil {
		left = maxPathSum(root.Left)
		nums = append(nums, left + root.Val)
	}
	if root.Right != nil {
		right = maxPathSum(root.Right)
		nums = append(nums, right + root.Val)
	}
	if root.Left != nil && root.Right != nil {
		nums = append(nums, left + right + root.Val)
	}

	sort.Ints(nums)

	nodeMax := nums[len(nums)-1]

	if root.Left != nil { nums = append(nums, left) }
	if root.Right != nil { nums = append(nums, right) }

	sort.Ints(nums)

	if res < nums[len(nums)-1] {
		res = nums[len(nums)-1]
	}

	if nodeMax < 0 { return 0 }
	return nodeMax
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
