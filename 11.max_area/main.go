package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1,2,4,3}))
}

// [1,2,4,3]
//  0 1 2 3
func maxArea(height []int) int {
	max, h, l := 0, 0, 0
	for i, j := 0, len(height)-1; i < j; {
		// 记录当前宽度
		l = j-i

		// 记录当前高度，同时谁高留谁
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
