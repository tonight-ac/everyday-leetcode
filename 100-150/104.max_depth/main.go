package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res int
func maxDepth(root *TreeNode) int {
	res = 0

	if root == nil { return res }

	recursion(root, 1)

	return res
}

func recursion(root *TreeNode, depth int) {
	if res < depth { res = depth }

	if root.Left != nil {
		recursion(root.Left, depth+1)
	}
	if root.Right != nil {
		recursion(root.Right, depth+1)
	}
}
