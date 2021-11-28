package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res [][]int
func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)

	if root == nil { return res }

	recursion(root, 0)

	return res
}

func recursion(n *TreeNode, depth int) {
	// 给res动态扩容
	for len(res) <= depth {
		res = append(res, make([]int, 0))
	}
	res[depth] = append(res[depth], n.Val)
	if n.Left != nil {
		recursion(n.Left, depth+1)
	}
	if n.Right != nil {
		recursion(n.Right, depth+1)
	}
}