package main

func isValidSudoku(board [][]byte) bool {
	row := make([][]bool, 9, 10)
	clo := make([][]bool, 9, 10)
	cube := make([][]bool, 9, 10)

	for i, r := range board {
		for j, v := range r {
			// 空格跳过
			if v == '.' {
				continue
			}

			if row[i][v-'0'] || clo[j][v-'0'] || cube[3*(i/3)+j/3][v-'0'] {
				return false
			}

			row[i][v-'0'], clo[j][v-'0'], cube[3*(i/3)+j/3][v-'0'] = true, true, true
		}
	}

	return true
}

// 这个版本速度比较慢，同时需要开过多的空间
// 并且需要在多个map中查询，我们尝试新的方法
//func isValidSudoku(board [][]byte) bool {
//	row := make([]map[byte]interface{}, 9)
//	clo := make([]map[byte]interface{}, 9)
//	cube := make([]map[byte]interface{}, 9)
//
//	// 如何通过下标计算cube位置
//	// i + j
//	// 4  7  10
//	// 7  9  12
//	// 10 12 15
//
//	// 3*(i/3) + j/3
//	// 0 1 2
//	// 3 4 5
//	// 6 7 8
//
//
//	// 初始化
//	for i := 0; i < 9; i++ {
//		row[i] = make(map[byte]interface{})
//		clo[i] = make(map[byte]interface{})
//		cube[i] = make(map[byte]interface{})
//	}
//
//	// 开始筛查
//	for i, r := range board {
//		for j, v := range r {
//			if v == '.' {
//				continue
//			}
//			// 对v更新并查重行、列、正方形
//			if _, ok := row[i][v]; ok {
//				return false
//			}
//			if _, ok := clo[j][v]; ok {
//				return false
//			}
//			if _, ok := cube[3*(i/3)+j/3][v]; ok {
//				return false
//			}
//			row[i][v], clo[j][v], cube[3*(i/3)+j/3][v] = nil, nil, nil
//		}
//	}
//
//	return true
//}
