package main

type ListNode struct {
	Val int
	Next *ListNode
}

// [1,2,3,3,4,4,5]
func main()  {
	list := []int{1,2,3,3,4,4,5}
	head := &ListNode{
		Val: 1,
	}

	for temp, i := head, 1; i < len(list); i++ {
		temp.Next = &ListNode{Val: list[i]}
		temp = temp.Next
	}
	deleteDuplicates(head)
}

// 以83题为基础
func deleteDuplicates(head *ListNode) *ListNode {
	// 加一个头节点，避免第一个也需要删除
	// -100 <= Node.val <= 100 所以设置为101
	head = &ListNode{Val: 101, Next: head}
	for cur := head; cur != nil && cur.Next != nil; {
		next := cur.Next // 下一个节点，这个节点也可能不安全
		forward := next.Next // 从下一个节点的下一个开始进行筛选
		for forward != nil && forward.Val == next.Val { forward = forward.Next }

		// 发生了移动，说明出现了重复
		if forward != next.Next {
			// 说明到末尾了
			cur.Next = forward
		} else { // 没有出现重复
			cur = cur.Next
		}
	}

	return head.Next
}
