package main

import "sort"

//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	//输入：nums = [-1,0,1,2,-1,-4]
	//输出：[[-1,-1,2],[-1,0,1]]

	// 先做排序
	sort.Ints(nums)

	return res
}
