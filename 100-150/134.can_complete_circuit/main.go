package main

import "fmt"

// 1, 2, 3, 4, 5
// 3, 4, 5, 1, 2
// -2 -2 -2 3  3

// 2, 3, 4
// 3, 4, 3
// -1 -1 1

//输入：
//[5,8,2,8]
//[6,5,6,6]
//输出：
//1
//预期结果：
//3
func main() {
	fmt.Println(canCompleteCircuit([]int{5,8,2,8}, []int{6,5,6,6}))
}

// 优化一下
// 主要利用的特性是，没有解，或者有解必唯一

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
	}

	if sum < 0 {
		return -1
	}

	currentGas := 0
	start := 0
	for i := 0; i < n; i++ {
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			currentGas = 0
			start = i + 1
		}
	}
	return start
}

// 先搞一版暴力
// 超时了nmd
//func canCompleteCircuit(gas []int, cost []int) int {
//	for i := 0; i < len(gas); i++ {
//		gas[i] -= cost[i]
//	}
//
//	for i := 0; i < len(gas); i++ {
//		tot := gas[i]
//		if tot < 0 { continue }
//		for j, k := (i + 1)%len(gas), 0; k < len(gas) - 1; k, j = k + 1, (j + 1)%len(gas) {
//			tot += gas[j]
//			if tot < 0 {
//				break
//			}
//		}
//		if tot >= 0 {
//			return i
//		}
//	}
//
//	return -1
//}

// -1 -11 -1 -1 13 -1 1
// 如果数值能够刚好持平或者更大，则要从首个最大正数的出发
//func canCompleteCircuit(gas []int, cost []int) int {
//	tot := 0
//	max, idx := 0, 0
//	for i := 0; i < len(gas); i++ {
//		gas[i] -= cost[i]
//		tot += gas[i]
//		if max < gas[i] {
//			max = gas[i]
//			idx = i
//		}
//	}
//
//	if tot < 0 { return -1 }
//	return idx
//}