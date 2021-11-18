package main

import "fmt"

//输入：
//[
//[1,2,3,4],
//[5,6,7,8],
//[9,10,11,12],
//[13,14,15,16]
//]
//输出：
//[1,2,3,4,8,12,16,15,14,13,9,10,11,7,6,5]
//预期结果：
//[1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
}

var dir = [][]int{0: {0, 1}, 1: {1, 0}, 2: {0, -1}, 3: {-1, 0}}
func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	res := make([]int, 0, n*m)

	// idx表示dir方向
	j, k, idx := 0, 0, 0
	// i表示结果集的第几位
	for i := 0; ; i++ {
		res = append(res, matrix[j][k])
		if len(res) == n*m { break }
		// matrix中的数字最大100，设置为101表示读取过了
		matrix[j][k] = 101
		// 按照右、下、左、上顺序更新j和k
		for l := idx; ; l++ {
			//fmt.Println(l%len(dir))
			// 避免越界
			v := dir[l%len(dir)]
			if !(j + v[0] >= 0 && j + v[0] < n) {
				continue
			}
			if !(k + v[1] >= 0 && k + v[1] < m) {
				continue
			}
			if matrix[j + v[0]][k + v[1]] == 101 {
				continue
			}
			j += v[0]
			k += v[1]
			idx = l
			break
		}
	}

	return res
}
