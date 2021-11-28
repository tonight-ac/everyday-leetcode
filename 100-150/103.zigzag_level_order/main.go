package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 102答案拿过来改一下
var res [][]int
func zigzagLevelOrder(root *TreeNode) [][]int {
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
	if depth%2 == 0 {
		res[depth] = append(res[depth], n.Val)
	} else {
		res[depth] = append(append(make([]int, 0), n.Val), res[depth]...)
	}

	if n.Left != nil {
		recursion(n.Left, depth+1)
	}
	if n.Right != nil {
		recursion(n.Right, depth+1)
	}
}