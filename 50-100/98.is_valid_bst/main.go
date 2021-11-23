package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return recursion(root, nil, nil)
}

// 左右两个数值判断
func recursion(n *TreeNode, left, right *int) bool {
	if n == nil { return true }
	if left != nil && n.Val <= *left { return false }
	if right != nil && n.Val >= *right { return false }
	return recursion(n.Left, left, &n.Val) && recursion(n.Right, &n.Val, right)
}