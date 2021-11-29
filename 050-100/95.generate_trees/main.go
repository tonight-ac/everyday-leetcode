package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 难点：1 如何遍历所有情况 2 如何巧妙的建多个不同的树
// 对于1 参考98题 对于2 先用数组存起来，最后再建树
//func main() {
	//t1 := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Right: &TreeNode{
	//			Val: 3,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 4,
	//	},
	//}
	//t2 := &TreeNode{}
	//copyTrees(t1, t2)
//}

func main() {
	fmt.Println(generateTrees(3))
}

// 递归太复杂了
// 1 <= n <= 8
var dp [9][9][]*TreeNode
func init() {
	for i := 1; i < len(dp); i++ { dp[i][i] = append(make([]*TreeNode, 0), &TreeNode{ Val: i }) }
}

func generateTrees(n int) []*TreeNode {
	return recursion(1, n)
}

func recursion(l, r int) []*TreeNode {
	if len(dp[l][r]) != 0 { return dp[l][r] }

	dp[l][r] = make([]*TreeNode, 0)

	for i := l; i <= r; i++ {
		if i - 1 < l { // 只有右侧
			for j := 0; j < len(recursion(i+1, r)); j++ {
				root := &TreeNode{ Val: i, Right: dp[i+1][r][j] }
				dp[l][r] = append(dp[l][r], root)
			}
		} else if r - i <= 0 { // 只有左侧
			for j := 0; j < len(recursion(l, i-1)); j++ {
				root := &TreeNode{ Val: i, Left: dp[l][i-1][j] }
				dp[l][r] = append(dp[l][r], root)
			}
		} else { // 左右都有
			for j := 0; j < len(recursion(l, i-1)); j++ {
				for k := 0; k < len(recursion(i+1, r)); k++ {
					root := &TreeNode{ Val: i, Left: dp[l][i-1][j], Right: dp[i+1][r][k] }
					dp[l][r] = append(dp[l][r], root)
				}
			}
		}
	}

	return dp[l][r]
}
// 1 <= n <= 8
// 先完成遍历
//var res []*TreeNode
//var uni [9]int
//var nn int
//func generateTrees(n int) []*TreeNode {
//	nn = n
//
//	uni = [9]int{}
//
//	res = make([]*TreeNode, 0)
//
//	for i := 1; i <= nn; i++ {
//		uni[i] = 1
//		recursion(&TreeNode{Val: i}, 0, nil, nil)
//		uni[i] = 0
//	}
//
//
//	return res
//}
//
//func recursion(head *TreeNode, count int, left, right *int) {
//	if count == nn {
//		temp := &TreeNode{}
//		copyTrees(head, temp)
//		res = append(res, temp)
//		return
//	}
//
//	start, end := 1, nn
//	if left != nil { start = *left+1 }
//	if right != nil { end = *right-1 }
//
//	for i := 1; i <= nn; i++ {
//		if uni[i] == 0 {
//			if i < head.Val { // 插入左边
//				head.Left = &TreeNode{ Val: i }
//				uni[i] = 1
//				recursion(head.Left, count+1)
//				uni[i] = 0
//				head.Left = nil
//			} else if i > head.Val {
//				head.Right = &TreeNode{ Val: i }
//				uni[i] = 1
//				recursion(head.Right, count+1)
//				uni[i] = 0
//				head.Right = nil
//			}
//		}
//	}
//}
//
//// 树的深拷贝
//func copyTrees(src *TreeNode, dst *TreeNode) {
//	if src == nil {
//		dst = nil
//		return
//	}
//
//	dst.Val = src.Val
//
//	if src.Left != nil {
//		dst.Left = &TreeNode{}
//		copyTrees(src.Left, dst.Left)
//	}
//
//	if src.Right != nil {
//		dst.Right = &TreeNode{}
//		copyTrees(src.Right, dst.Right)
//	}
//}
