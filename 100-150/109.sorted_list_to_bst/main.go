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
	left, mid, right := split(head)
	return &TreeNode{
		Val: mid.Val,
		Left: sortedListToBST(left),
		Right: sortedListToBST(right),
	}
}

func split(head *ListNode) (left *ListNode, mid *ListNode, right *ListNode){
	// 两个不好处理
	if head.Next.Next == nil { return nil, head, head.Next }

	prev, slow, quick := head, head, head.Next

	for quick != nil && quick.Next != nil {
		prev = slow
		slow = slow.Next
		quick = quick.Next.Next
	}

	// 对于left形成尾部
	prev.Next = nil

	return head, slow, slow.Next
}