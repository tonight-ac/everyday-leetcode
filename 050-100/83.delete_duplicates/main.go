package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 如果下一个节点跟当前数字相同，下一个节点删除
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }

	for cur := head; cur != nil && cur.Next != nil; {
		// 相同则删除next
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}