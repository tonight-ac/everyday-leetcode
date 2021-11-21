package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 如果下一个节点跟当前数字相同，下一节节点删除
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }

	for cur, next := head, head.Next; cur != nil && next != nil; {
		// 相同则删除next
		if cur.Val == next.Val {
			cur.Next = next.Next
			next = cur.Next
		} else {
			cur = cur.Next
			next = next.Next
		}
	}

	return head
}