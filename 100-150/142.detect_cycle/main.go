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

	//m := make(map[*ListNode]bool)
	//for head != nil {
	//	if m[head] == true {
	//		return head
	//	}
	//	m[head] = true
	//	head = head.Next
	//}

	return nil
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil { return false }
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast { return true }
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}