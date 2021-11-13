package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//  链表为空或只有一个，不需要反转
	if head == nil || head.Next == nil { return head }

	p := reverseList(head.Next)
	// head下一个保存的正是现在p开头的尾节点
	head.Next.Next = head
	// 制造head为尾巴
	head.Next = nil

	return p
}
