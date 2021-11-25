package main

import "fmt"

//{
//{'0','0','0','0','0','0','1'},
//{'0','0','0','0','1','1','1'},
//{'1','1','1','1','1','1','1'},
//{'0','0','0','1','1','1','1'}
//}

func main() {
	matrix := [][]byte{{'0','0','0','0','0','0','1'},{'0','0','0','0','1','1','1'},{'1','1','1','1','1','1','1'},{'0','0','0','1','1','1','1'}}
	//matrix := {}{}byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 { return 0 }
	n, m := len(matrix), len(matrix[0])

	// 动态数组 dp[i][j]表示matrix[i][j]和dp[i][j]/m和dp[i][j]%m能够构成最大矩形
	dp := make([][]int, n)
	for i := 0; i < n; i++ { dp[i] = make([]int, m) }

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' { continue }
			// 这里选取逻辑非常重要
			if i > 0 && j > 0 && matrix[i-1][j] == '1' && matrix[i][j-1] == '1' {
				x1, y1 := dp[i-1][j]/m, dp[i-1][j]%m
				for k := j; k >= y1; k-- {
					if matrix[i][k] == '0' {
						y1 = k+1
						break
					}
				}
				x2, y2 := dp[i][j-1]/m, dp[i][j-1]%m
				for k := i; k >= x2; k-- {
					if matrix[k][j] == '0' {
						x2 = k+1
						break
					}
				}
				// 计算哪个面积大
				if t1, t2 := (i-x1+1)*(j-y1+1), (i-x2+1)*(j-y2+1); t1 > t2 {
					dp[i][j] = x1*m+y1
				} else {
					dp[i][j] = x2*m+y2
				}
				//dp[i][j] = max(x1,x2)*m+max(y1,y2) // 取更靠近右下的
			} else if i > 0 && matrix[i-1][j] == '1' {
				x, _ := dp[i-1][j]/m, dp[i-1][j]%m
				dp[i][j] = x*m+j
			} else if j > 0 && matrix[i][j-1] == '1' {
				_, y := dp[i][j-1]/m, dp[i][j-1]%m
				dp[i][j] = i*m+y
			} else {
				dp[i][j] = i*m+j
			}
			if temp := (i-dp[i][j]/m+1)*(j-dp[i][j]%m+1); res < temp {
				res = temp
			}
		}
	}

	return res
}

//func max(a, b int) int {
//	if a > b { return a }
//	return b
//}
