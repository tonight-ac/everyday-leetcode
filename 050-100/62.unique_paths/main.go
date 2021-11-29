package main

// 动态规划不超时
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	// 第一行和第一列初始化为1
	for i := 0; i < m; i++ { matrix[i][0] = 1 }
	for j := 0; j < n; j++ { matrix[0][j] = 1 }

	// 开始DP
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[m-1][n-1]
}

// 暴力广度优先遍历，直接超时
//var mm, nn int
//var res int
//func uniquePaths(m int, n int) int {
//	mm, nn = m, n
//
//	res = 0
//
//	recursion(0, 0)
//
//	return res
//}
//
//func recursion(x, y int) {
//	if x == mm - 1 && y == nn - 1 {
//		res++
//		return
//	}
//
//	if x + 1 < mm {
//		recursion(x+1, y)
//	}
//
//	if y + 1 < nn {
//		recursion(x, y+1)
//	}
//}