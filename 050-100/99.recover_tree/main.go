package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//   1
// 2   3
// 左右两个数值判断

//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。

// 这个case的-33我感觉有很大问题
//输入：
//[146,71,-13,55,null,231,399,321,null,null,null,null,null,-33]
//输出：
//[71,321,231,55,null,146,399,-13,null,null,null,null,null,-33]
//预期结果：
//[146,71,321,55,null,231,399,-13,null,null,null,null,null,-33]

//        146
//     71     -13
//   55     231  399
// 321

//        146
//     71     321
//   55     231  399
// -13         -33
//输入：root = [3,1,4,null,null,2]
//输出：[2,1,4,null,null,3]
//解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。

//    3
//  1   2
//   4

//    4
//  1   2
//   3

//    3
//  1   4
//   2
func main() {
	//root := &TreeNode{
	//	Val: 3,
	//	Right: &TreeNode{
	//		Val: 2,
	//	},
	//	Left: &TreeNode{
	//		Val: 1,
	//		Right: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//}
	root := &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
		//Left: &TreeNode{
		//	Val: 3,
		//	Right: &TreeNode{
		//		Val: 2,
		//	},
		//},
	}
	recoverTree(root)
}

// 堆不行
// 先存储成堆的形式
// 修正错误
// 再恢复回去
//var temp []*TreeNode
//func recoverTree(root *TreeNode) {
//	// [1, 10^4]
//	temp = make([]*TreeNode, 10000)
//
//	// 把树打平成数组
//	treeToHeap(root, 0)
//
//	// 修复数据
//
//}
//
//func treeToHeap(root *TreeNode, idx int) {
//	if root == nil {
//		return
//	}
//	temp[idx] = root
//	treeToHeap(root.Left, idx*2+1)
//	treeToHeap(root.Right, idx*2+2)
//}
//
//func headToTree() {
//
//}
func recoverTree(root *TreeNode)  {
	var nums []int
	// 中序遍历，结果放进nums
	// 用中序遍历的原因，左子树必然在自己左边，右子树必然在自己右边
	// 而左边必然小于自己，右边的必然大于自己，寻找不符合升序要求的就行了
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	// 修正顺序 x y 为需要修复的数据
	x, y := findTwoSwapped(nums)
	// 复原
	recover(root, 2, x, y)
}
//    3
//  1   2
//   4

// 2 1 3 4
// 1 4 3 2
func findTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i + 1] < nums[i] {
			index2 = i + 1
			if index1 == -1 { // 先预设只有一个存在问题
				index1 = i
			} else { // 两个均已经找到
				break
			}
		}
	}
	// 做交换
	x, y := nums[index1], nums[index2]
	return x, y
}


func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	// 给数值做交换
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		// 需要修复的节点数归零
		if count == 0 {
			return
		}
	}
	recover(root.Right, count, x, y)
	recover(root.Left, count, x, y)
}

//func recursion(n, left, right *TreeNode) *TreeNode {
//	if n == nil { return nil }
//	var n1 *TreeNode
//	if n.Left != nil {
//		n1 = recursion(n.Left, left, n) // 左子树有问题的节点
//
//	}
//	var n2 *TreeNode
//	if n.Right != nil {
//		n2 = recursion(n.Right, n, right) // 右子树有问题的节点
//	}
//
//	if n1 == nil && n2 == nil {
//		if left != nil && n.Val < left.Val { return n }
//		if right != nil && n.Val > right.Val { return n }
//		return nil
//	}
//
//	// 左右子树都有一个有问题，直接交换，根据题目要求，直接交换1
//	if n1 != nil && n2 != nil {
//		n1.Val, n2.Val = n2.Val, n1.Val
//		return nil // 问题已经结局
//	}
//
//	if n1 != nil {
//		// 左边有问题，并且现在就能解决
//		if n1.Val > n.Val {
//			n1.Val, n.Val = n.Val, n1.Val
//			return nil
//		}
//		// 不能解决
//		return n1
//	}
//
//	if n2 != nil {
//		// 右边有问题，并且现在就能解决
//		if n2.Val < n.Val {
//			n2.Val, n.Val = n.Val, n2.Val
//			return nil
//		}
//		// 不能解决
//		return n2
//	}
//
//	return nil
//}

// 借鉴98题
// 先整理一下顺序
// 很遗憾又一次超时了
//func recoverTree(root *TreeNode)  {
//	recursion(root, nil, nil)
//}
//
//func recursion(n, left, right *TreeNode) {
//	if n == nil { return }
//
//	// 左儿树先清理一遍
//	if left != nil {
//		if n.Val < left.Val {
//			n.Val, left.Val = left.Val, n.Val
//		}
//	}
//	recursion(n.Left, left, n)
//	// 右子树再清理一遍
//	if right != nil {
//		if n.Val > right.Val {
//			n.Val, right.Val = right.Val, n.Val
//		}
//	}
//	recursion(n.Right, n, right)
//	// 先把当前节点和左节点调整
//	if n.Left != nil && n.Left.Val > n.Val { n.Val, n.Left.Val = n.Left.Val, n.Val }
//	// 再把当前节点和右节点调整
//	if n.Right != nil && n.Right.Val < n.Val { n.Val, n.Right.Val = n.Right.Val, n.Val }
//	// 左子树同时再清理一遍
//	if left != nil {
//		if n.Val < left.Val {
//			n.Val, left.Val = left.Val, n.Val
//		}
//	}
//	recursion(n.Left, left, n)
//	// 右子树同时再清理一遍
//	if right != nil {
//		if n.Val > right.Val {
//			n.Val, right.Val = right.Val, n.Val
//		}
//	}
//	recursion(n.Right, n, right)
//}
