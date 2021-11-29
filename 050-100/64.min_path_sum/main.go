package main

//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。

// 思路：选择左和上更小的和自己相加
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 1; i < m; i++ { grid[i][0] += grid[i-1][0] }
	for j := 1; j < n; j++ { grid[0][j] += grid[0][j-1] }

	// 开始DP
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
