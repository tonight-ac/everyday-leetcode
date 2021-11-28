package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// TODO 未完成
// 超时了
// 借用104
var balanced bool
func isBalanced(root *TreeNode) bool {
	balanced = true

	recursion(root)

	return balanced
}

func recursion(root *TreeNode) int {
	if root == nil { return 0 }
	if abs(recursion(root.Left) - recursion(root.Right)) > 1 { balanced = false }
	return max(recursion(root.Left), recursion(root.Right))+1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a > b { return a }
	return b
}
