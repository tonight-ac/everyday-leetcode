package main

import "math"

//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// -4 -1 1 2

func threeSumClosest(nums []int, target int) int {
	// 把前三个切下来 尝试一下贪心算法
	tuple := append(make([]int, 0), nums[0:3]...)
	res := tuple[0]+tuple[1]+tuple[2]
	diff := math.Abs(float64(res-target))

	// 从第四个开始遍历
	for i := 3; i < len(nums); i++ {
		//min := 0
		for j := 0; j < 3; j++ {
			if diff > math.Abs(float64(res - tuple[0] + nums[i] - target)) {

			}
		}
	}

	return res
}
