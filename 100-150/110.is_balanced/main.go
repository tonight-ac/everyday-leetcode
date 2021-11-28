package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 借用104
var balanced bool
func isBalanced(root *TreeNode) bool {
	balanced = true

	recursion(root)

	return balanced
}

func recursion(root *TreeNode) int {
	if root == nil { return 0 }
	left := recursion(root.Left)
	right := recursion(root.Right)
	if abs(left - right) > 1 { balanced = false }
	return max(left, right)+1
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
