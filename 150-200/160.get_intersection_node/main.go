package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 双指针法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB // 这部是关键，让A的指针指向了B链表
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA // 这部是关键，让B的指针指向了A链表 通过这种方式消弭了两个链表长度的差异
		} else {
			pb = pb.Next
		}
	}
	return pa
}
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	// 每次从headA中选一个节点，在整个headB中搜索
//	for headA != nil {
//		temp := headB
//		for temp != nil {
//			if temp == headA {
//				return headA
//			}
//			temp = temp.Next
//		}
//		headA = headA.Next
//	}
//
//	return nil
//}