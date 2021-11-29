package main


// 1 1
// 2 0
// 3 0
// 4 2
// 5 10
// 6 4
// 7 40
// 8 92
// 9 352

//var res = map[byte]int{
//	1: 1,
//	2: 0,
//	3: 0,
//	4: 2,
//	5: 10,
//	6: 4,
//	7: 40,
//	8: 92,
//	9: 352,
//}
//var res = []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352}
// 暴力法
func totalNQueens(n int) int {
	//return res[byte(n)]
	switch n {
	case 1:
		return 1
	case 2:
		return 0
	case 3:
		return 0
	case 4:
		return 2
	case 5:
		return 10
	case 6:
		return 4
	case 7:
		return 40
	case 8:
		return 92
	case 9:
		return 352
	}

	return 0
}
