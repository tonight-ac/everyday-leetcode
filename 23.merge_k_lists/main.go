package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 进行递归处理
func mergeKLists(lists []*ListNode) *ListNode {
	// 无链表，不需要处理
	if len(lists) == 0 { return nil }

	// 使用右移来做除2，提高效率
	return recursion(lists[:len(lists)>>1], lists[len(lists)>>1:])
}

func recursion(l1, l2 []*ListNode) *ListNode {
	var h1, h2 *ListNode

	if len(l1) > 1 {
		// 拆分两段递归
		h1 = recursion(l1[:len(l1)>>1], l1[len(l1)>>1:])
	} else if len(l1) == 1 { //
		h1 = l1[0]
	} // l1为空，则h1必为空，不需要赋值

	if len(l2) > 1 {
		// 拆分两段递归
		h2 = recursion(l2[:len(l2)>>1], l2[len(l2)>>1:])
	} else if len(l2) == 1 {
		h2 = l2[0]
	} // l2为空，则h2必为空，不需要赋值

	return mergeTwoLists(h1, h2)
}

// 复用21题的两个链表的归并问题
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
