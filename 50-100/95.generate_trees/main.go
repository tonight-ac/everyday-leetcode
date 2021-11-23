package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO 未完成
// 难点：1 如何遍历所有情况 2 如何巧妙的建多个不同的树
// 对于1 参考98题 对于2 先用数组存起来，最后再建树
func main() {
	fmt.Println(generateTrees(3))
}

// 1 <= n <= 8
// 先完成遍历
var nums = []int{1,2,3,4,5,6,7,8}
var uni [9]int
var nn int
func generateTrees(n int) []*TreeNode {
	nn = n

	uni = [9]int{}

	return nil
}

func recursion(head *TreeNode, left, right []int) []*TreeNode {
	lefts := make([]*TreeNode, 0)
	for i := 0; i < len(left); i++ {
		temp := &TreeNode{
			Val: head.Val,
			Left: &TreeNode{
				Val: left[i],
			},
		}
		lefts = append(lefts, recursion(temp.Left, left[:i], left[i+1:])...)
	}

	rights := make([]*TreeNode, 0)
	for i := 0; i < len(right); i++ {
		temp := &TreeNode{
			Val: head.Val,
			Right: &TreeNode{
				Val: right[i],
			},
		}
		rights = append(rights, recursion(temp.Right, right[:i], right[i+1:])...)
	}

	return nil
}

// 树粘贴
func copyTrees(src *TreeNode, dst *TreeNode) {

}
