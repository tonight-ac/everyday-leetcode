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

// TODO 未完成

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

func recoverTree(root *TreeNode)  {
	_ = recursion(root, nil, nil)
}

func recursion(n, left, right *TreeNode) *TreeNode {
	if n == nil { return nil }
	var n1 *TreeNode
	if n.Left != nil {
		n1 = recursion(n.Left, left, n) // 左子树有问题的节点

	}
	var n2 *TreeNode
	if n.Right != nil {
		n2 = recursion(n.Right, n, right) // 右子树有问题的节点
	}

	if n1 == nil && n2 == nil {
		if left != nil && n.Val < left.Val { return n }
		if right != nil && n.Val > right.Val { return n }
		return nil
	}

	// 左右子树都有一个有问题，直接交换，根据题目要求，直接交换1
	if n1 != nil && n2 != nil {
		n1.Val, n2.Val = n2.Val, n1.Val
		return nil // 问题已经结局
	}

	if n1 != nil {
		// 左边有问题，并且现在就能解决
		if n1.Val > n.Val {
			n1.Val, n.Val = n.Val, n1.Val
			return nil
		}
		// 不能解决
		return n1
	}

	if n2 != nil {
		// 右边有问题，并且现在就能解决
		if n2.Val < n.Val {
			n2.Val, n.Val = n.Val, n2.Val
			return nil
		}
		// 不能解决
		return n2
	}

	return nil
}

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
