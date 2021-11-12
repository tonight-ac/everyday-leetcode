package main

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// 长度小于等于1，换无可换
	if head == nil || head.Next == nil {
		return head
	}

	// 加一个头节点，处理起来更方便
	node := &ListNode{
		Next: head,
	}
	head = node

	for node.Next != nil && node.Next.Next != nil {
		// 下一个和下两个的节点先记录下来
		// 1. 0 -> 1 -> 2 -> 3
		one := node.Next
		two := node.Next.Next

		// 2. 0 -> 2 -> 3
		//     1 -^
		node.Next = two

		// 3. 0 -> 2 -> 3
		//         1 -^
		one.Next = two.Next

		// 4. 0 -> 2 -> 1 -> 3
		two.Next = one

		// 移动node跳过两个开始下一轮遍历
		node = node.Next.Next
	}

	return head.Next
}
