package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { return false }
	if targetSum == 0 { return true }
	if targetSum >= root.Val {
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	}
	return false
}