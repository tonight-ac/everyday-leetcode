package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//Input: preorder = [3, 9,1, 20,15,7], inorder = [1,9, 3, 15,20,7]
//Output: [3,9,20,null,null,15,7]
//    3
//  9  20
// 1  15 7
// 递归构建
// 分组找规律
//preorder = [3, 9,1, 20,15,7], inorder = [1,9, 3, 15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { return nil }
	if len(preorder) == 1 { return &TreeNode{Val: preorder[0]} }
	idx := 0
	for idx < len(inorder) && inorder[idx] != preorder[0] { idx++ }
	return &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1:idx+1], inorder[0:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}
