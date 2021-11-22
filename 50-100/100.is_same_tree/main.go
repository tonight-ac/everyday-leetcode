package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归解决
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}
