package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 借用112
var res [][]int
func pathSum(root *TreeNode, targetSum int) [][]int {
	res = make([][]int, 0)

	recursion(root, targetSum, make([]int, 0))
	return res
}

func recursion(root *TreeNode, targetSum int, one []int) {
	if root == nil { return }
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		one = append(one, root.Val)
		res = append(res, append(make([]int, 0), one...))
		return
	}
	recursion(root.Left, targetSum-root.Val, append(one, root.Val))
	recursion(root.Right, targetSum-root.Val, append(one, root.Val))
}