package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(postorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
}

// 1
//  3
// 2

// 迭代算法
func postorderTraversal(root *TreeNode) []int {
	if root == nil { return nil }

	in := append(make([]*TreeNode, 0), root)

	res := make([]int, 0)

	for {
		var first *TreeNode
		if len(in) != 0 {
			first = in[len(in)-1]
		} else {
			break
		}

		if first.Left == nil && first.Right == nil {
			in = in[:len(in)-1]
			res = append(res, first.Val)
		}
		if first.Right != nil {
			in = append(in, first.Right)
			first.Right = nil
		}
		if first.Left != nil {
			in = append(in, first.Left)
			first.Left = nil
		}
	}

	return res
}

// 递归算法
// 消耗内存但优美版
//func postorderTraversal(root *TreeNode) []int {
//	if root == nil { return nil }
//	left := postorderTraversal(root.Left)
//	right := postorderTraversal(root.Right)
//	return append(left, append(right, root.Val)...)
//}
