package main

type ListNode struct {
	Val int
	Next *ListNode
}

//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//L0 → L1 → … → Ln - 1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 1 2 3 // 4 5
// 1 5 2 4 3

// 1 2 3 4 // 5 6
// 1 6 2 5 3 4
func reorderList(head *ListNode) {
	// 先从中间截断，拆为两个链表，然后将后半段链表反转，再插入

	// 双指针判断位置

	// 反转后半段 借用206题

	// 插入


}

func recursion() {

}