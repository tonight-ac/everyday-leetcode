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

	connectThem([]*Node{root})

	return root
}

// 换一种方式
func connectThem(them []*Node) {
	if len(them) == 0 { return }
	next := make([]*Node, 0)
	for i := 0; i < len(them); i++ {
		if i + 1 < len(them) {
			them[i].Next = them[i+1]
		}
		if them[i].Left != nil { next = append(next, them[i].Left) }
		if them[i].Right != nil { next = append(next, them[i].Right) }
	}
	connectThem(next)
}