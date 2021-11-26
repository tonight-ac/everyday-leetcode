package main

import "fmt"



func main() {
	//matrix := [][]byte{{'0','0','0','0','0','0','1'},{'0','0','0','0','1','1','1'},{'1','1','1','1','1','1','1'},{'0','0','0','1','1','1','1'}}
	matrix := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	fmt.Println(maximalRectangle(matrix))
}

// 在搜索的基础上做一下剪枝
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 { return 0 }

	n, m := len(matrix), len(matrix[0])
	size := n*m

	res := 0

	for i := 0; i < size; i++ {
		x1, x2 := i/m, i%m
		// 0不能做左上角
		if matrix[x1][x2] == '0' { continue }
		dp := make([]bool, size)
		for y1 := x1; y1 < n; y1++ {
			for y2 := x2; y2 < m; y2++ {
				// 0不能做右下角
				if matrix[y1][y2] == '0' { continue }
				j := y1*m+y2
				if x1 == y1 && x2 == y2 {
					// 相同节点，判断是否等于1
					dp[j] = true
				} else if x1 == y1 {
					// 同一行
					dp[j] = dp[j-1]
				} else if x2 == y2 {
					// 同一列
					dp[j] = dp[j-m]
				} else {
					// 左边和上边都是可以的
					dp[j] = dp[j-1] && dp[j-m]
					if !dp[i] { break } //
				}
				if temp := (y1-x1+1)*(y2-x2+1); dp[j] && res < temp {
					res = temp
				}
			}
		}
	}

	return res
}

//// 这不是dp，这是一个妥妥的搜索
//func maximalRectangle(matrix [][]byte) int {
//	if len(matrix) == 0 || len(matrix[0]) == 0 { return 0 }
//
//	n, m := len(matrix), len(matrix[0])
//	size := n*m
//
//	res := 0
//
//	for i := 0; i < size; i++ {
//		x1, x2 := i/m, i%m
//		// 0不能做左上角
//		if matrix[x1][x2] == '0' { continue }
//		dp := make([]bool, size)
//		for j := i; j < size; j++ {
//			y1, y2 := j/m, j%m
//			// 矩形面积需要为正
//			if y2 < x2 { continue }
//			// 0不能做右下角
//			if matrix[y1][y2] == '0' { continue }
//			if x1 == y1 && x2 == y2 {
//				// 相同节点，判断是否等于1
//				dp[j] = true
//			} else if x1 == y1 {
//				// 同一行
//				dp[j] = dp[j-1]
//			} else if x2 == y2 {
//				// 同一列
//				dp[j] = dp[j-m]
//			} else {
//				// 左边和上边都是可以的
//				dp[j] = dp[j-1] && dp[j-m]
//				if !dp[i] { break } //
//			}
//			if temp := (y1-x1+1)*(y2-x2+1); dp[j] && res < temp {
//				res = temp
//			}
//		}
//	}
//
//	return res
//}

// 这个方法在下面这个case中失效了
// 不能堵漏洞，换一个思路
//{
//{'0','0','0','0','0','0','1'},
//{'0','0','0','0','1','1','1'},
//{'1','1','1','1','1','1','1'},
//{'0','0','0','1','1','1','1'}
//}
//func maximalRectangle(matrix [][]byte) int {
//	if len(matrix) == 0 || len(matrix[0]) == 0 { return 0 }
//	n, m := len(matrix), len(matrix[0])
//
//	// 动态数组 dp[i][j]表示matrix[i][j]和dp[i][j]/m和dp[i][j]%m能够构成最大矩形
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ { dp[i] = make([]int, m) }
//
//	res := 0
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			if matrix[i][j] == '0' { continue }
//			// 这里选取逻辑非常重要
//			if i > 0 && j > 0 && matrix[i-1][j] == '1' && matrix[i][j-1] == '1' {
//				x1, y1 := dp[i-1][j]/m, dp[i-1][j]%m
//				for k := j; k >= y1; k-- {
//					if matrix[i][k] == '0' {
//						y1 = k+1
//						break
//					}
//				}
//				x2, y2 := dp[i][j-1]/m, dp[i][j-1]%m
//				for k := i; k >= x2; k-- {
//					if matrix[k][j] == '0' {
//						x2 = k+1
//						break
//					}
//				}
//				// 计算哪个面积大
//				if t1, t2 := (i-x1+1)*(j-y1+1), (i-x2+1)*(j-y2+1); t1 > t2 {
//					dp[i][j] = x1*m+y1
//				} else {
//					dp[i][j] = x2*m+y2
//				}
//				//dp[i][j] = max(x1,x2)*m+max(y1,y2) // 取更靠近右下的
//			} else if i > 0 && matrix[i-1][j] == '1' {
//				x, _ := dp[i-1][j]/m, dp[i-1][j]%m
//				dp[i][j] = x*m+j
//			} else if j > 0 && matrix[i][j-1] == '1' {
//				_, y := dp[i][j-1]/m, dp[i][j-1]%m
//				dp[i][j] = i*m+y
//			} else {
//				dp[i][j] = i*m+j
//			}
//			if temp := (i-dp[i][j]/m+1)*(j-dp[i][j]%m+1); res < temp {
//				res = temp
//			}
//		}
//	}
//
//	return res
//}

//func max(a, b int) int {
//	if a > b { return a }
//	return b
//}
