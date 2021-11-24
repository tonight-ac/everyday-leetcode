package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO 未完成
func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	recoverTree(root)
	fmt.Println(1)
}


//   1
// 2   3
// 左右两个数值判断

//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。

//输入：
//[146,71,-13,55,null,231,399,321,null,null,null,null,null,-33]
//输出：
//[71,321,231,55,null,146,399,-13,null,null,null,null,null,-33]
//预期结果：
//[146,71,321,55,null,231,399,-13,null,null,null,null,null,-33]

//输入：root = [3,1,4,null,null,2]
//输出：[2,1,4,null,null,3]
//解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。

// 借鉴98题
func recoverTree(root *TreeNode)  {
	recursion(root, nil, nil)
}

func recursion(n, left, right *TreeNode) {
	if n == nil { return }
	// 先把当前节点和左节点调整
	if n.Left != nil && n.Left.Val > n.Val { n.Val, n.Left.Val = n.Left.Val, n.Val }
	// 再把当前节点和右节点调整
	if n.Right != nil && n.Right.Val < n.Val { n.Val, n.Right.Val = n.Right.Val, n.Val }
	if left != nil {
		if n.Val < left.Val {
			n.Val, left.Val = left.Val, n.Val
		}
	}
	if right != nil {
		if n.Val > right.Val {
			n.Val, right.Val = right.Val, n.Val
		}
	}
	recursion(n.Left, left, n)
	recursion(n.Right, n, right)
}
