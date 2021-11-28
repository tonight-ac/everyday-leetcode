package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil { return false }
	return recursion(root.Left, root.Right)
}

func recursion(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l != nil && r != nil {
		if l.Val != r.Val { return false }
	} else {
		return false
	}

	return recursion(l.Left, r.Right)  && recursion(l.Right, r.Left)
}