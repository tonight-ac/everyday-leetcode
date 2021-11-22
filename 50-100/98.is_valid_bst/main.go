package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO
func isValidBST(root *TreeNode) bool {
	if root == nil { return true }

	return recursion(root.Left, root.Val, true, false) && recursion(root.Right, root.Val, false, true)
}

func recursion(root *TreeNode, std int, left, right bool) bool {
	if root == nil { return true }
	if left && root.Val < std {
		return recursion(root.Left, root.Val, true, false) && recursion(root.Right, root.Val, false, true)
	}
	if right && root.Val > std {
		return recursion(root.Left, root.Val, true, false) && recursion(root.Right, root.Val, false, true)
	}
	return false
}