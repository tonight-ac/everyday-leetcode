package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归算法
// 消耗内存但优美版
func postorderTraversal(root *TreeNode) []int {
	if root == nil { return nil }
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	return append(left, append(right, root.Val)...)
}
