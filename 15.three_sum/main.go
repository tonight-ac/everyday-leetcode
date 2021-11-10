package main

import "sort"

//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	//输入：nums = [-1,0,1,2,-1,-4]
	//输出：[[-1,-1,2],[-1,0,1]]

	// 先做排序，不排序的话，无法进行，且无法去重
	sort.Ints(nums)

	// 首尾两个指针，同时i < j+1需要留出三个数字的空间
	for i, j := 0, len(nums)-1; i < j+1; i, j = i+1, j-1 {
		// 求二数之和取反，然后在i+1, j-1的范围内二分找sum
		sum := -(nums[i] + nums[j])

		// 二分法
		for l, r := i + 1, j - 1; l <= r; {
			mid := (l + r) / 2
			if sum < nums[mid] {
				r = mid - 1
			} else if sum > nums[mid] {
				l = mid + 1
			} else {
				// 确实找到了，加入结果集合
				res = append(res, []int{nums[i], sum, nums[j]})
				break
			}
		}

		// 去除重复元素
		for i < j + 1 && nums[i] == nums[i+1] {
			i++
		}
		// 去除重复元素
		for i < j + 1 && nums[j] == nums[j-1] {
			j--
		}
	}

	return res
}
