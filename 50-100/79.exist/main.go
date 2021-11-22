package main

// TODO 未完成
var n, m int
var b [][]byte
var uni []int

//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true

var dir = [][]int{0: {0, 1}, 1: {1, 0}, 2: {0, -1}, 3: {-1, 0}}

func exist(board [][]byte, word string) bool {
	b = board

	n, m = len(board), len(board[0])

	uni = make([]int, n*m)

	recursion(0, 0, word)

	return false
}

func recursion(x, y int, s string) bool {
	if len(s) > 0 && b[x][y] == s[0] { s = s[1:] }

	// s匹配完毕
	if len(s) == 0 { return true }

	// 越界了
	//if idx >= n*m { return false }
	//
	//x, y := idx/m, idx%m

	//flag := false
	for i := 0; i < len(dir);i++ {
		xx, yy := x+dir[i][0], y+dir[i][1]
		if xx >= 0 && xx < n && yy >= 0 && yy < m {

		}
	}

	return false
}
