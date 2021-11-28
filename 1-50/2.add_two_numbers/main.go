package main


type ListNode struct {
	Val int
	Next *ListNode
}

// 字符串加法 本质：模拟加法

// 不好处理的地方在于进位

// 目标：节约内存 能不申请不新申请 复用已存在的节点

// 1 2 3
// 2 3 4

// 5 7 9
// 4 5 6 7 8 9 1

// 6 8 0 4
// 5 6 7

// 2个十进制数
// 最多进几位
// 9999 9999
// xxxxx
//
//
// 三步走
// 1 求公共长部分的和
// 2 拼接多余需要处理的节点到l1
// 3 处理进位

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

	// 将多余部分全部拼接到l1
	if l1.Next == nil {
		l1.Next = l2.Next
	}

	// 处理进位
	for {
		if div == 0 {
			break
		}

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