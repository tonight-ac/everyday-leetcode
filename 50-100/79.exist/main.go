package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED"))
}

//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true

// TODO 未完成
var n, m int
var b [][]byte
var uni []int
var length int
var dir = [][]int{0: {0, 1}, 1: {1, 0}, 2: {0, -1}, 3: {-1, 0}}

func exist(board [][]byte, word string) bool {
	b = board

	n, m = len(board), len(board[0])

	length = len(word)

	uni = make([]int, n*m)

	recursion(0, word)

	return false
}
// 当前匹配，向下递归，不匹配直接扫下一个
func recursion(idx int, s string) bool {
	// s匹配完毕
	if len(s) == 0 { return true }

	// 越界了
	if idx >= n*m { return false }

	x, y := idx/m, idx%m

	flag := false
	if b[x][y] != s[0] {
		if len(s) == length {
			// 继续向下找，首个匹配
			flag = recursion(idx+1, s)
		} else {
			return false
		}
	} else {
		uni[idx] = 1
		// 当前匹配
		for i := 0; i < len(dir);i++ {
			xx, yy := x+dir[i][0], y+dir[i][1]
			if xx >= 0 && xx < n && yy >= 0 && yy < m {
				if uni[xx*m+yy] == 0 {
					flag = recursion(xx*m+yy, s[1:])
					if flag {
						break
					}
				}
			}
		}
		uni[idx] = 0
	}

	return flag
}
