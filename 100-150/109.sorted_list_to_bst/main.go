package main

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO
// 借鉴108 通过快慢指针，判断中间位置
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil { return nil }
	if head.Next == nil { return &TreeNode{ Val: head.Val }}
	slow, quick := head, head
	// 1 2 3 4 5
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return nil
}