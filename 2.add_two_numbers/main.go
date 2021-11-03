package main


type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tot, mod, div := 0, 0, 0
	head := l1

	// 先处理公共部分
	for {
		tot = l1.Val + l2.Val + div
		mod, div = tot%10, tot/10

		l1.Val = mod

		if l1.Next == nil || l2.Next == nil {
			break
		}

		l1, l2 = l1.Next, l2.Next
	}

	// l2可能没有，也可能有
	// 总之将l2多余部分全部拼接到l1
	if l1.Next == nil {
		l1.Next = l2.Next
	}

	// 从拼接处开始处理进位
	for {
		// 没有进位，无须处理
		if div == 0 {
			break
		}

		// 剩最后一次进位了
		if l1.Next == nil {
			l1.Next = &ListNode{ Val: 1 }
			break
		}

		l1 = l1.Next

		tot = l1.Val + div
		mod, div = tot%10, tot/10

		l1.Val = mod
	}

	return head
}