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

// 借鉴98题
func recoverTree(root *TreeNode)  {
	recursion(root, nil, nil)
}

// 左右两个数值判断
func recursion(n, left, right *TreeNode) {
	if n == nil { return }
	if n.Left != nil && n.Left.Val > n.Val { n.Val, n.Left.Val = n.Left.Val, n.Val }
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
