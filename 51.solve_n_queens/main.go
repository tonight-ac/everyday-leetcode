package main

import (
	"fmt"
	"strings"
)

// 递归搜索
// 0 行 1 列 2 左斜 3 右斜

// 左斜为 (i-j)+(n-1)
// 右斜为 i+j
// 0,0 0,1 0,2 0,3
// 1,0 1,1 1,2 1,3
// 2,0 2,1 2,2 2,3
// 3,0 3,1 3,2 3,3

func main() {
	fmt.Println(len(solveNQueens(9)))

}

var uni [4][18]byte
var res [][]string
var length int
func solveNQueens(n int) [][]string {
	uni = [4][18]byte{}

	res = make([][]string, 0)

	length = n

	recursion(0, n, make([]string, 0))

	return res
}

func recursion(idx , n int, s []string) {
	// x和y都迭代到右下角时
	if idx >= length*length {
		if n == 0 {
			temp := make([]string, 0)
			for i := 0; i < len(s); i+=length {
				temp = append(temp, strings.Join(s[i:i+length], ""))
			}
			res = append(res, temp)
		}
		return
	}

	x, y := idx/length, idx%length

	row, clo, left, right := x, y, (x-y)+(length-1), x+y
	if uni[0][row] + uni[1][clo] + uni[2][left] + uni[3][right] == 0 {
		s = append(s, "Q")
		uni[0][row], uni[1][clo], uni[2][left], uni[3][right] = 1, 1, 1, 1
		recursion(idx+1, n - 1, s)
		uni[0][row], uni[1][clo], uni[2][left], uni[3][right] = 0, 0, 0, 0
		s = s[:len(s)-1]
	}
	s = append(s, ".")
	recursion(idx+1, n, s)
	s = s[:len(s)-1]
}
