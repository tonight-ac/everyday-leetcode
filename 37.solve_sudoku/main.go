package main

import "fmt"

func main() {
	n := 80
	fmt.Println(n/9, n%9)
}

var row, clo, cube [9][10]byte
func solveSudoku(board [][]byte) {
	row, clo, cube = [9][10]byte{}, [9][10]byte{}, [9][10]byte{}

	// 把已经存在的加入map中方便查重
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
			row[i][n], clo[j][n], cube[3*(i/3)+j/3][n] = 0, 0, 0
		}
	} else { // 有数字直接跳过，不需要填写
		recursion(board, idx + 1)
	}
}

//输入：
//[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
//输出：
//[["5","3","4","6","7","8","9","9","8"],["6","7","7","1","9","5","7","3","4"],["2","9","8","3","3","4","7","6","7"],["8","5","9","9","6","7","9","9","3"],["4","5","9","8","5","3","9","9","1"],["7","5","9","9","2","1","8","4","6"],["9","6","9","5","3","7","2","8","4"],["3","8","7","4","1","9","6","3","5"],["3","4","4","6","8","2","6","7","9"]]
//预期结果：
//[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]