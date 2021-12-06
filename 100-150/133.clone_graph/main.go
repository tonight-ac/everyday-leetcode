package main

import "fmt"

type Node struct {
	Val int
	Neighbors []*Node
}

func main() {
	n1 := &Node{Val: 1, Neighbors: make([]*Node, 0)}
	n2 := &Node{Val: 2, Neighbors: make([]*Node, 0)}
	n3 := &Node{Val: 3, Neighbors: make([]*Node, 0)}
	n4 := &Node{Val: 4, Neighbors: make([]*Node, 0)}
	n1.Neighbors = append(n1.Neighbors, n2, n4)
	n2.Neighbors = append(n2.Neighbors, n1, n3)
	n3.Neighbors = append(n3.Neighbors, n2, n4)
	n4.Neighbors = append(n4.Neighbors, n1, n3)
	fmt.Println((cloneGraph(n1).Neighbors)[0])
}
var mapping map[*Node]*Node
func cloneGraph(node *Node) *Node {
	mapping = make(map[*Node]*Node)
	return recursion(node)
}

func recursion(node *Node) *Node {
	if node == nil { return nil }
	if mapping[node] != nil { return mapping[node] }
	n := &Node{ Val: node.Val, Neighbors: make([]*Node, 0) }
	mapping[node] = n
	for _, v := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, recursion(v))
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
