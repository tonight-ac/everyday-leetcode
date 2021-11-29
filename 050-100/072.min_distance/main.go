package main

import "fmt"

//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符

//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// 这题经典，但是真的不会
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	// 有一个字符串为空串
	if n * m == 0 {
		return n + m
	}

	// DP 数组
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, m+1)
	}

	// 边界状态初始化
	for i := 0; i < n + 1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m + 1; j++ {
		dp[0][j] = j
	}

	// 计算所有 DP 值
	for i := 1; i < n + 1; i++ {
		for j := 1; j < m + 1; j++ {
			left := dp[i - 1][j] + 1
			down := dp[i][j - 1] + 1
			leftDown := dp[i - 1][j - 1]
			if word1[i - 1] != word2[j - 1] {
				leftDown += 1
			}
			dp[i][j] = min(left, min(down, leftDown))
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

