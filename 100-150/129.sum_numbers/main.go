package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res int
func sumNumbers(root *TreeNode) int {
	res = 0

	recursion(root, 0)

	return res
}

func recursion(root *TreeNode, n int) {
	if root.Left == nil && root.Right == nil {
		res += n * 10 + root.Val
		return
	}
	if root.Left != nil {
		recursion(root.Left, n * 10 + root.Val)
	}
	if root.Right != nil {
		recursion(root.Right, n * 10 + root.Val)
	}
}