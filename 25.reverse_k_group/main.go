package main

type ListNode struct {
	Val int
	Next *ListNode
}

//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 常见手段，加一个头节点
	head = &ListNode{ Next: head }

	// 找到开始位置，反转链表
	for start, end, cnt := head, head, 0; end != nil; end, cnt = end.Next, cnt + 1 {
		if cnt == k {
			// 截断返回
			//
			temp := end.Next
			end.Next = nil
			_ = reverseList(start.Next)
			end = start.Next
			end.Next = temp

			// start.Next目前是尾巴了
			start.Next.Next = temp

			// end目前头部了
			start.Next = end

			// 更新cnt进行下一个计数周期
			cnt = 0
		}

	}

	return head.Next
}

// 借用206题题解，反转链表
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


