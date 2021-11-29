package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//    3
//  9  20
// 1  15 7

//[2,1,4,null,null,3]
//输出：
//[2,null,1,null,4,3]
//预期结果：
//[2,null,1,null,4,null,3]
// [3, 9,1, 20,15,7]
// O(1)空间法
// 把左右儿子树旋转，如果右子树不为空，将右子树缀在交换过的左子树最右下角
func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) { return }
	if root.Left == nil {
		flatten(root.Right)
	} else if root.Right == nil {
		// 左放到右节点
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
	} else {
		temp := root.Right
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
		for root.Right != nil {
			root = root.Right
		}
		root.Right = temp
		flatten(root.Right)
	}
}

// O(n)空间法
//var temp []*TreeNode
//func flatten(root *TreeNode) {
//	// 先打平为前序遍历
//	temp = make([]*TreeNode, 0)
//
//	inorder(root)
//
//	for i := 0; i < len(temp); i++ {
//		temp[i].Left = nil
//		if i != len(temp) - 1 {
//			temp[i].Right = temp[i+1]
//		}
//	}
//}
//
//func inorder(root *TreeNode) {
//	if root == nil { return }
//	temp = append(temp, root)
//	inorder(root.Left)
//	inorder(root.Right)
//}
