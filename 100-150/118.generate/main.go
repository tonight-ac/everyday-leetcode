package main

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	res[0] = []int{1}

	for i := 2; i <= numRows; i++ {
		res[i] = make([]int, i)
		for j := 0; j < len(res[i]); j++ {
			if j > 0 {
				res[i][j] += res[i-1][j-1]
			}
			if j < len(res[i-1]) {
				res[i][j] += res[i-1][j]
			}
		}
	}

	return res
}
