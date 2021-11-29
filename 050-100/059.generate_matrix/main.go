package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

// 54题的逆过程，其实是一样的因为遍历过程是一样的
var dir = [][]int{0: {0, 1}, 1: {1, 0}, 2: {0, -1}, 3: {-1, 0}}
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	// 开空间，初始化结果集
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	// idx表示dir方向
	j, k, idx := 0, 0, 0
	for i := 1; ; i++ {
		//fmt.Println(i)
		matrix[j][k] = i
		if i == n*n { break }
		// 按照右、下、左、上顺序更新j和k
		for l := idx; ; l++ {
			// 避免越界
			v := dir[l%len(dir)]
			if !(j + v[0] >= 0 && j + v[0] < n) {
				continue
			}
			if !(k + v[1] >= 0 && k + v[1] < n) {
				continue
			}
			if matrix[j + v[0]][k + v[1]] != 0 {
				continue
			}
			j += v[0]
			k += v[1]
			idx = l
			break
		}
	}

	return matrix
}
