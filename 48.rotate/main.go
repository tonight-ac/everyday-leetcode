package main

import "fmt"

// 1.
//[1,2,3] 待填入 2 3
//[4,5,6]
//[7,8,9]

// 2.
//[1,2,3] 待填入 6 9
//[4,5,2]
//[7,8,3]

// 3.
//[1,2,3] 待填入 8 7
//[4,5,2]
//[9,6,3]

// 4.
//[7,2,3] 待填入 4 1
//[8,5,2]
//[9,6,3]

// 5.
//[7,4,1] 完毕
//[8,5,2]
//[9,6,3]

//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

func main() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rotate(matrix)
	//printM(matrix)
}

// 原地旋转填入法
func rotate(matrix [][]int) {
	n := len(matrix)
	last := n - 1

	// 只需要处理n/2即可
	for i := 0; i < n/2; i++ {
		// 缓冲数组
		list := make([]*int, 0)

		// 每次两边均缩短一距离
		for j := i+1; j < n-i; j++ {
			list = append(list, &matrix[i][j])
		}

		// 执行上图中第2步操作
		for j := i+1; j < n-i; j++ {
			list = append(list, &matrix[j][last-i]) // 为之后作准备
			matrix[j][last-i] = *list[0] // 移动到目标位置
			list = list[1:] // 将第一个删除
		}

		// 执行上图中第3步操作
		//for j := last-i; j >= i; j-- {
		//
		//}

		// 执行上图中第4步操作

		// 执行上图中第5步操作
	}
}

// swap 原位置交换
func swap(matrix [][]int, i1, j1, i2, j2 int) {
	temp := matrix[i1][j1]
	matrix[i1][j1] = matrix[i2][j2]
	matrix[i2][j2] = temp
}

func printM(matrix [][]int) {
	for _, row := range matrix {
		for _, v := range row {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func printL(list []*int) {
	for _, v := range list {
		fmt.Print(*v, " ")
	}
	fmt.Println()
}
