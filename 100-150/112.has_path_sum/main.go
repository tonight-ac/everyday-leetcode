package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//输入：
//[5,4,8,11,null,13,4,7,2,null,null,null,1]
//22
//输出：
//false
//预期结果：
//true
// TODO
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { return false }
	if targetSum == 0 { return true }
	if targetSum >= root.Val {
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	}
	return false
}