package main

type ListNode struct {
	Val int
	Next *ListNode
}

//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]

// 1 2 3 4 5
// 截取1 2
// 反转 再拼接回去
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 常见手段，加一个头节点
	head = &ListNode{ Next: head }

	// 找到开始位置，反转链表
	for start, end, cnt := head, head, 0; end != nil; end, cnt = end.Next, cnt + 1 {
		if cnt == k {
			// 截断返回
			// 记录下3
			nextFirst := end.Next
			end.Next = nil
			// 归位end
			end = start.Next

			// 反转1 -> 2
			thisFirst := reverseList(start.Next)

			// 把反转后的链表拼接回去
			// start.Next已经反转到末尾，接上 1 -> 3
			end.Next = nextFirst
			// 连接上0 —> 2
			start.Next = thisFirst
			// 把start也归位
			start = end

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
