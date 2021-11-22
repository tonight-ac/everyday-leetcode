package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO 未完成
func main() {
	fmt.Println(generateTrees(3))
}

//  1 2 3
var nn int
var uni []int
var res []*TreeNode
func generateTrees(n int) []*TreeNode {
	nn = n

	uni = make([]int, nn+1)

	recursion(nil, &TreeNode{})

	return res
}

func recursion(n *TreeNode, head *TreeNode) {
	if len(uni) == nn {
		res = append(res, head)
		return
	}
	for i := 1; i <= nn; i++ {
		if uni[i] != 1 {
			uni[i] = 1
			if head.Val == 0 {
				head.Val = i
				recursion(head, head)
			} else if i < n.Val {
				n.Left = &TreeNode{ Val: i }
				recursion(n.Left, head)
				n.Left = nil
			} else {
				n.Right = &TreeNode{ Val: i }
				recursion(n.Right, head)
				n.Right = nil
			}
			uni[i] = 0
		}
	}
}