package main

// 做多33行
var res [34]int
func getRow(rowIndex int) []int {

	return res[:rowIndex+1]
}
