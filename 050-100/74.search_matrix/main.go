package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1},{3}}, 3))
}

// 先按照对行首列进行二分搜索
// 在对行进行二分搜索
func searchMatrix(matrix [][]int, target int) bool {
	// 先对第一列进行二分搜索确定行
	l, r, mid := 0, len(matrix) - 1, 0
	for ; l <= r; {
		mid = (l + r) >> 1
		if target >= matrix[mid][0] {
			if mid == len(matrix)-1 || target < matrix[mid+1][0] {
				break
			} else {
				l = mid + 1
			}
		} else {
			r = mid - 1
		}
	}

	list, l, r := matrix[mid], 0, len(matrix[mid]) - 1
	// 确定行之后在这行进行二分搜索
	for ; l <= r; {
		mid := (l + r) >> 1
		if target < list[mid] {
			r = mid - 1
		} else if target > list[mid] {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
