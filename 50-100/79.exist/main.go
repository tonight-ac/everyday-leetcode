package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED"))
}
//[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
//"SEE"
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true

var n, m int
var b [][]byte
var uni []bool
var dir = [][]int{0: {0, 1}, 1: {1, 0}, 2: {0, -1}, 3: {-1, 0}}

func exist(board [][]byte, word string) bool {
	b = board

	n, m = len(board), len(board[0])

	uni = make([]bool, n*m)

	// 从每个点开始广度优先搜索
	for i := 0; i < n*m; i++ {
		if ok := recursion(i, word); ok {
			return true
		}
	}

	return false
}
// 当前匹配，向下递归，不匹配直接扫下一个
func recursion(idx int, s string) bool {
	uni[idx] = true
	defer func() { uni[idx] = false }()

	x, y := idx/m, idx%m

	if s[0] != b[x][y] { return false }

	s = s[1:]

	// s匹配完毕
	if len(s) == 0 { return true }

	for _, d := range dir {
		xx, yy := x + d[0], y + d[1]
		if xx >= 0 && xx < n && yy >= 0 && yy < m {
			next := xx*m+yy
			if !uni[next] {
				uni[next] = true
				if ok := recursion(next, s); ok {
					return true
				}
				uni[next] = false
			}
		}
	}

	return false
}
