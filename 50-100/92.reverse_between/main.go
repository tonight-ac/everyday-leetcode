package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 需要用到19的指针和206题的反转数组
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right { return head }

	return nil
}
