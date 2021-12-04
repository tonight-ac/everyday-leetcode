package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	head.Next.Next = head
	fmt.Println(detectCycle(head))
}

// 借用141题
// 先判断环，然后"切"环
// 再将题目转化为，寻找两个链表的交叉点
// 相交链表参照160题
func detectCycle(head *ListNode) *ListNode {
	cross := hasCycle(head)
	if cross == nil { return nil } // 不存在环

	return getIntersectionNode(head, cross.Next, cross)
}

// 160题的改造版本，双指针寻找最早交叉点
func getIntersectionNode(headA, headB, end *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == end {
			pa = headB // 这部是关键，让A的指针指向了B链表
		} else {
			pa = pa.Next
		}
		if pb == end {
			pb = headA // 这部是关键，让B的指针指向了A链表 通过这种方式消弭了两个链表长度的差异
		} else {
			pb = pb.Next
		}
	}
	return pa
}

// 141题的改造版本
func hasCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return nil }
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast { return slow }
		slow = slow.Next
		fast = fast.Next.Next
	}

	return nil
}