package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1,2,4,3}))
}

// [1,2,4,3]
//  0 1 2 3

// [1,8,6,2,5,4,8,3,7]
//  0 1 2 3 4 5 6 7 8
//      i           j

// 2 3 2 1
// i   j

// 9 8 1 1 1 1 1

// 短板效应：面积不由高的决定，但会由矮的限制
func maxArea(height []int) int {
	max, h, l := 0, 0, 0
	for i, j := 0, len(height)-1; i < j; {
		// 0啥也装不了，跳过
		if height[i] == 0 {
			i++
			continue
		}
		if height[j] == 0 {
			j--
			continue
		}
		// 记录当前宽度
		l = j-i

		// 记录当前高度，同时谁高留谁
		// 移动谁，都在宽度上必然有损失，保留更长，避免成为之后可能的短板

		// [1,100,6,2,5,4,8,200,7]
		//      i               j

		//       i             j
		//     i            j
		if height[i] <= height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}

		if temp := l*h; max < temp {
			max = temp
		}
	}

	return max
}
