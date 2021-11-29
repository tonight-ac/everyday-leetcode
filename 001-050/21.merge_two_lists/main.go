package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 均为空，则结果必为空
	if l1 == nil && l2 == nil { return nil }

	// 加一个头节点，处理起来更方便
	head := &ListNode{}
	node := head

	// 两个链表长度可能不登长，先处理公共长部分
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1 // 添加l1
			l1 = l1.Next
		} else {
			node.Next = l2 // 添加l2
			l2 = l2.Next
		}
		node = node.Next
	}

	// 最后处理可能超出的部分
	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}

	return head.Next
}
