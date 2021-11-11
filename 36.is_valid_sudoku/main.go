package main

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]interface{}, 9)
	clo := make([]map[byte]interface{}, 9)
	cube := make([]map[byte]interface{}, 9)

	// i + j
	// 4  7  10
	// 7  9  12
	// 10 12 15

	// 3*(i/3) + j/3
	// 0 1 2
	// 3 4 5
	// 6 7 8


	// 初始化
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]interface{}, 9)
		clo[i] = make(map[byte]interface{}, 9)
		cube[i] = make(map[byte]interface{}, 9)
	}

	// 开始筛查
	for i, r := range board {
		for j, v := range r {
			// 对v更新并查重行、列、正方形
			if _, ok := row[i][v]; ok {
				return false
			}
			if _, ok := clo[j][v]; ok {
				return false
			}
			if _, ok := cube[3*(i/3)+j/3][v]; ok {
				return false
			}
			row[i][v], clo[j][v], cube[i+j][v] = nil, nil, nil
		}
	}

	return true
}
