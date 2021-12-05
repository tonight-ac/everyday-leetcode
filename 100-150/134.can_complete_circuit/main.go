package main

import "fmt"

// 1, 2, 3, 4, 5
// 3, 4, 5, 1, 2
// -2 -2 -2 3  3

// 2, 3, 4
// 3, 4, 3
// -1 -1 1

func main() {
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3}))
}

// -1 -11 -1 -1 13 -1 1
// 如果数值能够刚好持平或者更大，则要从首个最大正数的出发
func canCompleteCircuit(gas []int, cost []int) int {
	tot := 0
	max, idx := 0, 0
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
		tot += gas[i]
		if max < gas[i] {
			max = gas[i]
			idx = i
		}
	}

	if tot < 0 { return -1 }
	return idx
}