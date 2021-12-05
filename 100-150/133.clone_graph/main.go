package main

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil { return nil }
	if len(node.Neighbors) == 0 { return &Node{ Val: node.Val, Neighbors: node.Neighbors }}
	n := &Node{ Val: node.Val, Neighbors: make([]*Node, 0) }
	for _, v := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, cloneGraph(v))
	}
	return n
}
