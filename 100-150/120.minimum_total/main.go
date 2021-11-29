package main

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j-1]
			} else if j == len(triangle[i-1]) - 1 {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}

	res := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if res > triangle[len(triangle)-1][i] {
			res = triangle[len(triangle)-1][i]
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}