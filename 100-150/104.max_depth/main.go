package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 另一种办法
func maxDepth(root *TreeNode) int {
	if root == nil { return 0 }
	return max(maxDepth(root.Left), maxDepth(root.Right))+1
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

// 其中一种办法
var res int
func recursion(root *TreeNode, depth int) {
	if res < depth { res = depth }

	if root.Left != nil {
		recursion(root.Left, depth+1)
	}
	if root.Right != nil {
		recursion(root.Right, depth+1)
	}
}

