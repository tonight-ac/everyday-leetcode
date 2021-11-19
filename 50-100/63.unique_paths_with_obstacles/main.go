package main

import "fmt"

//输入：
//[
//[0,0,0],
//[0,1,0],
//[0,0,0]]
//输出：
//1
//预期结果：
//2
// 基于62题做一些改动，为了节约空间，在obstacleGrid上原地修改

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0,0,0},{0,1,0,},{0,0,0}}))
}

//输入：
//[[1,0]]
//输出：
//1
//预期结果：
//0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 { return 0 }

	m, n := len(obstacleGrid), len(obstacleGrid[0])

	if m == 1 && n == 1 { return 1-obstacleGrid[0][0] }

	// 第一行和第一列初始化为1
	for i, b := 1, 1; i < m; i++ {
		// 如果存在障碍，这一行剩下的都是不可达的
		if obstacleGrid[i][0] == 1 { b = 0 }
		obstacleGrid[i][0] = b
	}
	for j, b := 1, 1; j < n; j++ {
		// 如果存在障碍，这一列剩下的都是不可达的
		if obstacleGrid[0][j] == 1 { b = 0 }
		obstacleGrid[0][j] = b
	}

	// 开始DP
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
