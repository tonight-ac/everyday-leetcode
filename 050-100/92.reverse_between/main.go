package main

type ListNode struct {
	Val int
	Next *ListNode
}

// [1,2,3,3,4,4,5]
func main()  {
	list := []int{1,2,3,4,5}
	head := &ListNode{
		Val: 1,
	}

	for temp, i := head, 1; i < len(list); i++ {
		temp.Next = &ListNode{Val: list[i]}
		temp = temp.Next
	}
	reverseBetween(head, 2, 4)
}

// 需要用到206题的反转数组
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right { return head }

	// 加一个头节点，方便处理头部
	head = &ListNode{ Next: head }

	leftNode, rightNode := head, head

	// 对于1 2 3 4 5 leftNode需要指向1（提前一个，方便处理） rightNode指向4
	// 定位leftNode
	for ; left > 1; left-- { leftNode = leftNode.Next }
	// 定位rightNode
	for ; right > 0; right-- { rightNode = rightNode.Next }

	// 人为制造终止条件，方便反转数组时判断末尾
	follow := rightNode.Next
	rightNode.Next = nil
	reverseList(leftNode.Next)
	leftNode.Next.Next = follow
	leftNode.Next = rightNode

	return head.Next
}

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