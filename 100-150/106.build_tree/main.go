package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//中序遍历 inorder = [9,1, 3, 15,20,7]
//后序遍历 postorder = [1,9, 15,7,20, 3]
//返回如下的二叉树：
//     3
//  9    20
//   1  15  7
// postorder的末尾必然为当前节点
// 借鉴105题目的分组方式
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 { return nil }
	if len(postorder) == 1 { return &TreeNode{Val: postorder[len(postorder)-1]} }
	idx := 0
	for idx < len(inorder) && inorder[idx] != postorder[len(postorder)-1] { idx++ }
	return &TreeNode{
		Val: postorder[len(postorder)-1],
		Left: buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
}
