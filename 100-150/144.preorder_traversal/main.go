package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(preorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
}

// 递归算法
// 消耗内存但优美版
func preorderTraversal(root *TreeNode) []int {
	if root == nil { return nil }
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	return append(append(append(make([]int, 0), root.Val), left...), right...)
}