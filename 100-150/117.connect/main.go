package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

// 参考116
func connect(root *Node) *Node {
	if root == nil { return nil }

	connectTwo(root.Left, root.Right)

	return root
}

func connectTwo(one, two *Node) {
	if one != nil {
		one.Next = two
		connectTwo(one.Left, one.Right)
	}

	if two != nil {
		connectTwo(two.Left, two.Right)
	}

	if one != nil && two != nil {
		connectTwo(one.Right, two.Left)
	}
}