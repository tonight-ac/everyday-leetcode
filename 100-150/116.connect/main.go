package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil { return nil }

	connectTwo(root.Left, root.Right)

	return root
}

func connectTwo(one, two *Node) {
	if one == nil || two == nil { return }
	one.Next = two
	connectTwo(one.Left, one.Right)
	connectTwo(two.Left, two.Right)
	connectTwo(one.Right, two.Left)
}