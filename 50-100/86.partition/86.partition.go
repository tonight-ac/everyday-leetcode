package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	list := []int{1}
	head := &ListNode{
		Val: 1,
	}

	for temp, i := head, 1; i < len(list); i++ {
		temp.Next = &ListNode{Val: list[i]}
		temp = temp.Next
	}
	partition(head, 0)
}

// 小于x的放一列，大于等于x的放另外一列
// 然后再把两列拼接在一起
func partition(head *ListNode, x int) *ListNode {
	sHead, lHead := &ListNode{}, &ListNode{}
	sEnd, lEnd := sHead, lHead

	for ; head != nil; {
		if head.Val < x { // 小的放进s
			sEnd.Next = head
			sEnd = sEnd.Next
			head = head.Next
			sEnd.Next = nil
		} else { // 大的放进l
			lEnd.Next = head
			lEnd = lEnd.Next
			head = head.Next
			lEnd.Next = nil
		}
	}

	// s没有的话直接返回l
	if sHead.Next == nil {
		return lHead.Next
	}

	// 把l拼接到s后面
	sEnd.Next = lHead.Next

	// 返回s
	return sHead.Next
}
