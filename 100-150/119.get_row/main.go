package main

// 做多33行
var res [34]int
func init() {
	res[0] = 1
}
//[
//[1] 0
//[1 1] 1
//[1 2 1] 2     0. (i-1)+(i-1) 2*i-1-(i-2)
//[1 3 3 1] 3   0. (i-1)+(i-2)
//[1 4 6 4 1] 4 0. (i-1)+(i-3) (i-1)+(i-1) 2*i-1-1*1
//[1 5 10 10 5 1] 5 i (i-1)+(i-2)+(i-2) 3*i-1-2*2
//]

// 先填前rowIndex/2+1 后面的对称复制
func getRow(rowIndex int) []int {
	// 先填前rowIndex/2+1
	i := 0
	for ; i < rowIndex/2+1; i++ {
		// 根据i判断当前下标数值，这是一个找规律题目

	}

	// 对称复制
	for ; i <= rowIndex; i++ {
		res[i] = res[rowIndex-i] // TODO
	}

	return res[:rowIndex+1]
}
