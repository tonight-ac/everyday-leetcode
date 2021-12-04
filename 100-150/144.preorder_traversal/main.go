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
	//fmt.Println(preorderTraversal(&TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 1,
	//		Right: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}))
}

// 迭代算法
func preorderTraversal(root *TreeNode) []int {
	if root == nil { return nil }
	left := make([]*TreeNode, 0)
	right := make([]*TreeNode, 0)
	res := make([]int, 0)

	for {
		if len(left) != 0 {
			root = left[0]
			left = left[1:]
		} else if len(right) != 0 {
			root = right[0]
			right = right[1:]
		} else if root == nil {
			break
		}

		res = append(res, root.Val)
		if root.Left != nil { left = append(left, root.Left) }
		if root.Right != nil { right = append(append(make([]*TreeNode, 0), root.Right), right...) }
		root = nil
	}

	return res
}

// 递归算法
// 消耗内存但优美版
//func preorderTraversal(root *TreeNode) []int {
//	if root == nil { return nil }
//	left := preorderTraversal(root.Left)
//	right := preorderTraversal(root.Right)
//	return append(append(append(make([]int, 0), root.Val), left...), right...)
//}