package main

// 技巧性太强，不看答案很难想到
func trap(height []int) int {
	// 短板效应
	// 判断一个格子能不能放下水，取决于自身高度和两边的最大高度
	if len(height) == 0 { return 0 }
	length := len(height)

	lMax := make([]int, length)
	lMax[0] = height[0]
	rMax := make([]int, length)
	rMax[length - 1] = height[length - 1]

	for i := 1; i < length; i++ {
		lMax[i] = max(height[i], lMax[i - 1])
	}

	for i := length - 2; i >= 0; i-- {
		rMax[i] = max(height[i], rMax[i + 1]);
	}

	res := 0
	for i := 1; i < length - 1; i++ {
		res += min(lMax[i], rMax[i]) - height[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}