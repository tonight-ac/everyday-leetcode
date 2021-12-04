package main

import "fmt"

func main() {
	board := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
	solve(board)
	fmt.Println(board)
}

// 先从四周有'O'的位置出发去搜索，将能搜索到的全部用'A'标记
// 然后再遍历，将'A'改为'O'，将'O'改为x
var b [][]byte
var m, n int
var dir = [][]int{{1, 0},{0, 1},{-1, 0}, {0, -1}}
func solve(board [][]byte) {
	m, n = len(board), len(board[0])

	if m == 1 || n == 1 { return }

	b = board

	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' { recursion(i, 1) }
		if board[i][n-1] == 'O' { recursion(i, n-2) }
	}

	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' { recursion(1, i) }
		if board[m-1][i] == 'O' { recursion(m-2, i) }
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func recursion(i, j int) {
	if !(i > 0 && i < m-1 && j > 0 && j < n-1) { return }
	if b[i][j] == 'X' || b[i][j] == 'A' { return }
	b[i][j] = 'A'
	for k := 0; k < len(dir); k++ {
		x := i+dir[k][0]
		y := j+dir[k][1]
		//if x > 0 && x < m-1 && y > 0 && y < n-1 {
			recursion(x, y)
		//}
	}
}