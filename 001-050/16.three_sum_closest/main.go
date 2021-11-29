package main

import (
	"fmt"
	"sort"
)

//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// -4 -1 1 2

func main() {
	fmt.Println(threeSumClosest([]int{-55,-24,-18,-11,-7,-3,4,5,6,9,11,23,33}, 0))
}
// -10

//[-55,-24,-18,-11,-7,-3,4,5,6,9,11,23,33]
//0
//输出：
//2
//预期结果：
//0

//[1,2,4,8,16,32,64,128]
//82
//输出：
//88
//预期结果：
//82

func threeSumClosest(nums []int, target int) int {
	// 对数组进行排序
	sort.Ints(nums)

	// 默认结果为前三个的计算和
	res := nums[0] + nums[1] + nums[2]

	// 进行遍历
	for i := 0; i < len(nums); i++ {
		// 选定i之后，利用有序性从两边开始选择
		for start, end := i + 1, len(nums) - 1; start < end; {
			sum := nums[start] + nums[end] + nums[i]

			// 更新更接近的值
			if abs(target - sum) < abs(target - res) {
				res = sum
			}

			if sum > target {
				// 升序排列，选择更小的
				end--
			} else if sum < target {
				// 升序排列，选择更大的
				start++
			} else {
				// 直接命中返回
				return res
			}

		}
	}
	return res
}

// go 缺乏原生int的abs
func abs(x int) int {
	if x < 0 { return -x }
	return x
}

// 这里使用贪心，无法解决的问题在于，最接近的是0，存在刚好相等的情况
//func threeSumClosest(nums []int, target int) int {
//	sort.Ints(nums)
//
//	// 把前三个切下来 尝试一下贪心算法
//	tuple := append(make([]int, 0), nums[len(nums)-3:]...)
//	res := tuple[0] + tuple[1] + tuple[2]
//	diff := abs(res - target)
//
//	// 从第四个开始遍历
//	for i := len(nums)-4; i >= 0; i-- {
//		idx := -1
//		min := -1
//		for j := 0; j < len(tuple); j++ {
//			if temp := abs(res - tuple[j] + nums[i] - target); temp < diff {
//				// 初始化min 或者取更小的min
//				if min < 0 || min < diff - temp {
//					min = diff - temp
//					idx = j
//				}
//			}
//		}
//		if idx != -1 {
//			// swap
//			tuple[idx] = nums[i]
//			// 计算 res diff
//			res = tuple[0] + tuple[1] + tuple[2]
//			diff = abs(res - target)
//		}
//	}
//
//	return res
//}
//

