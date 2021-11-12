package main

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// 类似19题，通过快慢指针定位之后，做反转
	// 区别是k可能大于ListNode总长度
	// 所以先计算总长度，然后取mod

	// 不需要移动
	if k == 0 {
		return head
	}

	// 先计算链表长度
	length := 0
	for temp := head; temp != nil; temp = temp.Next {
		length++
	}

	// 对k取mod
	k = k%length
	// 如果刚好移动的长度是链表长度蒸熟倍，也不需要移动
	if k == 0 {
		return head
	}

	// 开始双指针定位

	return nil
}
