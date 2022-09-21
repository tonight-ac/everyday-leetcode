package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 最起码是O(n)，因为无法知道链表长度
// 可以通过双指针的方式 第一个指针先走n步
// 初态 指针1先行n步
// ----------------- m单位
// --------- n单位 指针1
// 指针2

// 终态 指针1走到末尾
// ----------------- m单位
// ---------++++++++ n单位"-" m-n单位"+" 指针1
// +++++++++-------- m-n单位"-" n单位"+" 指针2

// ++++++++--------- 一个节点走n个长度
// -----------------
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 为了方便操作，加一个空的head节点，可以轻松解决去除第一个的情况
	head = &ListNode{ Next: head }

	// x 1 2 3 4 5
	//          l1
	//      l2
	// 初始化两个指针
	l1, l2 := head, head

	// 移动l1 n个单位
	for ; n > 0; n-- { l1 = l1.Next }

	// l1移动到尾部，l2同步移动
	for ; l1.Next != nil; {
		l1 = l1.Next
		l2 = l2.Next
	}

	// 去除l2.Next节点
	l2.Next = l2.Next.Next

	return head.Next
}