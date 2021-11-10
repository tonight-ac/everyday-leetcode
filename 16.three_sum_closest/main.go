package main

//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// -4 -1 1 2

func threeSumClosest(nums []int, target int) int {
	// 把前三个切下来 尝试一下贪心算法
	tuple := append(make([]int, 0), nums[0:3]...)
	res := tuple[0] + tuple[1] + tuple[2]
	diff := abs(res - target)

	// 从第四个开始遍历
	for i := 3; i < len(nums); i++ {
		idx := -1
		min := -1
		for j := 0; j < len(tuple); j++ {
			if temp := abs(res - tuple[j] + nums[i] - target); temp < diff {
				// 初始化min 或者取更小的min
				if min < 0 || min > diff - temp {
					min = diff - temp
					idx = j
				}
			}
		}

		if idx != -1 {
			// swap
			tuple[idx] = nums[i]
			// 计算 res diff
			res = tuple[0] + tuple[1] + tuple[2]
			diff = abs(res - target)
		}
	}

	return res
}

// go 缺乏原生int的abs
func abs(x int) int {
	if x < 0 { return -x }
	return x
}
