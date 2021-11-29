package main

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil { return head }
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
	// 如果刚好移动的长度是链表长度整数倍，也不需要移动
	if k == 0 {
		return head
	}

	// 开始双指针定位，把倒数k个元素，放到开头位置
	// 初始化两个指针
	l1, l2 := head, head

	// 移动l1 n个单位
	for ; k > 0; k-- { l1 = l1.Next }

	// l1移动到尾部，l2同步移动
	for ; l1.Next != nil; {
		l1 = l1.Next
		l2 = l2.Next
	}

	// 重新粘连
	res := l2.Next
	l2.Next = nil // 制造一个尾部
	l1.Next = head // 连接头部

	return res
}
