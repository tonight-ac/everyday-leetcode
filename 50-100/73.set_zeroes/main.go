package main

func main() {
	m := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	setZeroes(m)
}

// TODO 寻找更优的常数空间算法
// 常数空间的做法
// 递归解法，遇到0递归将整行设置为0，如果递归过程中遇到0继续递归
//func setZeroes(matrix [][]int) {
//	m, n := len(matrix), len(matrix[0])
//
//	rZero, cZero := false, false
//
//	// 第一遍直接遇到0
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if matrix[i][j] == 0 {
//				// 先向左向右回溯置0，这个没有历史负担，之前已经等0的不需要设置0
//				for k := i-1; k >= 0 && matrix[k][j] != 0; k-- { matrix[k][j] = 0 }
//				for k := j-1; k >= 0 && matrix[i][k] != 0; k-- { matrix[i][k] = 0 }
//				// 处理列
//				if i+1 < m && matrix[i+1][j] != 0 {
//					rZero = false
//				}
//			}
//		}
//	}
//}

// O(n)空间解法
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	clo := make([]bool, n)

	for i := 0; i < m; i++ {
		zero := false
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				// 向前置0
				for k := j-1; k >= 0 && matrix[i][k] != 0; k-- { matrix[i][k] = 0 }
				// 向后置0
				zero = true

				clo[j] = true
			} else if zero {
				matrix[i][j] = 0
			}
		}
	}

	// 给列置0
	for j := 0; j < n; j++ {
		if clo[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}


// O(m + n)空间解法，先记录出需要置0的行号和列号，再进行更新
//func setZeroes(matrix [][]int) {
//	m, n := len(matrix), len(matrix[0])
//	row := make([]bool, m)
//	clo := make([]bool, n)
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if matrix[i][j] == 0 {
//				row[i] = true
//				clo[j] = true
//			}
//		}
//	}
//
//	// 给行置0
//	for i := 0; i < m; i++ {
//		if row[i] {
//			for j := 0; j < n; j++ {
//				matrix[i][j] = 0
//			}
//		}
//	}
//
//	// 给列置0
//	for j := 0; j < n; j++ {
//		if clo[j] {
//			for i := 0; i < m; i++ {
//				matrix[i][j] = 0
//			}
//		}
//	}
//}
