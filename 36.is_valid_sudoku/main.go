package main

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]interface{}, 9)
	clo := make([]map[byte]interface{}, 9)
	cube := make([]map[byte]interface{}, 9)

	for i, r := range board {
		for j, v := range r {
			// 对v同时更新三个区间
			row[i][v] = nil
			clo[j][v] = nil
			cube[i+j][v] = nil
		}
	}
}

