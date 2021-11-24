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
	if left != nil {
		if n.Val < left.Val {
			n.Val, left.Val = left.Val, n.Val
			//return
		}
	}
	recursion(n.Left, left, n)
	if left != nil {
		if n.Val < left.Val {
			n.Val, left.Val = left.Val, n.Val
			//return
		}
	}
	if right != nil {
		if n.Val > right.Val {
			n.Val, right.Val = right.Val, n.Val
			//return
		}
	}
	recursion(n.Right, n, right)
	if right != nil {
		if n.Val > right.Val {
			n.Val, right.Val = right.Val, n.Val
			//return
		}
	}

}
