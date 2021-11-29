package main

import "fmt"

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。

func main() {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}

// 这是人能想出来的吗？
// 从暴力法能不能衍生出什么思考？
func largestRectangleArea(heights []int) int {
	// 这里为了代码简便，在柱体数组的头和尾加了两个高度为 0 的柱体。
	temp := append(make([]int, 0), 0)
	temp = append(temp, heights...)
	temp = append(temp, 0)

	stack := make([]int, 0)
	res := 0
	for i := 0; i < len(temp); i++ {
		// 对栈中柱体来说，栈中的下一个柱体就是其「左边第一个小于自身的柱体」；
		// 若当前柱体 i 的高度小于栈顶柱体的高度，说明 i 是栈顶柱体的「右边第一个小于栈顶柱体的柱体」。
		// 因此以栈顶柱体为高的矩形的左右宽度边界就确定了，可以计算面积🌶️ ～
		for len(stack) != 0 && temp[i] < temp[stack[len(stack)-1]] {
			h := temp[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if res < (i - stack[len(stack)-1] - 1) * h {
				res = (i - stack[len(stack)-1] - 1) * h
			}
		}
		stack = append(stack, i)
	}

	return res
}
// 跟42题是不是有相似性，或者第11题
// 暴力法，不用想一个超时
//func largestRectangleArea(heights []int) int {
//	//
//	res := 0
//	for i := 0; i < len(heights); i++ {
//		r := i
//		for j := i; j < len(heights); j++ {
//			if heights[j] < heights[i] {
//				break
//			}
//			r = j
//		}
//		l := i
//		for j := i; j >= 0; j-- {
//			if heights[j] < heights[i] {
//				break
//			}
//			l = j
//		}
//		if res < (r-l+1)*heights[i] {
//			res = (r-l+1)*heights[i]
//		}
//	}
//	return res
//}
