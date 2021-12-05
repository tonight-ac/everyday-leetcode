package main

type Node struct {
	Val int
	Neighbors []*Node
}

var uni map[*Node]bool
func cloneGraph(node *Node) *Node {
	uni = make(map[*Node]bool)
	return nil
}

func recursion(node *Node) *Node {
	if node == nil { return nil }
	if uni[node] { return nil }
	if len(node.Neighbors) == 0 { return &Node{ Val: node.Val, Neighbors: node.Neighbors }}
	n := &Node{ Val: node.Val, Neighbors: make([]*Node, 0) }
	for _, v := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, cloneGraph(v))
	}
	return n
}

// 内存溢出了，这是无向 连通 图
//func cloneGraph(node *Node) *Node {
//	if node == nil { return nil }
//	if len(node.Neighbors) == 0 { return &Node{ Val: node.Val, Neighbors: node.Neighbors }}
//	n := &Node{ Val: node.Val, Neighbors: make([]*Node, 0) }
//	for _, v := range node.Neighbors {
//		n.Neighbors = append(n.Neighbors, cloneGraph(v))
//	}
//	return n
//}
