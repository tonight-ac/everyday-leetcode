package main

import "fmt"

func main() {
	n := 80
	fmt.Println(n/9, n%9)
}

var row, clo, cube [9][10]byte
func solveSudoku(board [][]byte) {
	row, clo, cube = [9][10]byte{}, [9][10]byte{}, [9][10]byte{}

	for i, r := range board {
		for j, v := range r {
			if v == '.' { continue }
			nums := v-'0'
			row[i][nums], clo[j][nums], cube[3*(i/3)+j/3][nums] = 1, 1, 1
		}
	}

	recursion(board, 0)
}
// 8/9 = 0
// 0 1 2 3 4 5 6 7 8
// 9 10 11 12 13 14 15 16 17
// ...
func recursion(board [][]byte, idx int) {
	// 结束条件，算上0一共81个格子
	if idx == 80 {
		return
	}

	// i，j通过计算得到
	i, j := idx / 9, idx % 9

	// . 表示需要待填入
	if board[i][j] == '.' {
		// 尝试所有可能
		for n := 1; n <= 9; n++ {
			if row[i][n] + clo[j][n] + cube[3*(i/3)+j/3][n] > 0 {
				continue
			}
			board[i][j] = byte(n) + '0'
			row[i][n], clo[j][n], cube[3*(i/3)+j/3][n] = 1, 1, 1
			recursion(board, idx + 1)
			//board[i][j] = '.' 不需要清空
			row[i][n], clo[j][n], cube[3*(i/3)+j/3][n] = 0, 0, 0
		}
	} else { // 有数字直接跳过，不需要填写
		recursion(board, idx + 1)
	}
}