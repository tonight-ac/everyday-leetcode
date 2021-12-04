package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{100,4,200,1,3,2}))
}

// 时间复杂度为 O(n) 解法
func longestConsecutive(nums []int) int {
	if len(nums) == 0 { return 0 }
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	max := 1
	for i := 0; i < len(nums); i++ {
		left, right := nums[i]-1, nums[i]+1
		// 先求右边界
		for m[right] {
			m[right] = false
			right++
		}
		// 再求左边界
		for m[left] {
			m[left] = false
			left--
		}
		m[nums[i]] = false
		if max < right - left - 1 {
			max = right - left - 1
		}
	}

	return max
}

// 先使用最简单的排序法
//func longestConsecutive(nums []int) int {
//	if len(nums) == 0 { return 0 }
//	sort.Ints(nums)
//	max, count := 1, 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i] == nums[i-1]+1 {
//			count++
//		} else if nums[i] != nums[i-1] {
//			count = 1
//		}
//		if max < count {
//			max = count
//		}
//	}
//
//	return max
//}

//func longestConsecutive(nums []int) int {
//	m := make(map[int]int)
//	count := make([][]int, 0)
//	for i := 0; i < len(nums); i++ {
//		if idx, ok := m[nums[i]]; ok {
//			delete(m, nums[i])
//			if nums[i] <= count[idx][0] {
//				count[idx][0] = nums[i]
//				m[nums[i]-1] = idx
//			} else {
//				count[idx][1] = nums[i]
//				m[nums[i]+1] = idx
//			}
//		} else {
//			count = append(count, []int{nums[i], nums[i]})
//			m[nums[i]-1] = len(count)-1
//			m[nums[i]+1] = len(count)-1
//		}
//	}
//
//	res := 1
//	for i := 0; i < len(count); i++ {
//		if res < count[i][1] - count[i][0] + 1 {
//			res = count[i][1] - count[i][0] + 1
//		}
//	}
//
//	return res
//}

//func longestConsecutive(nums []int) int {
	//m := make(map[int]int)
	//for i := 0; i < len(nums); i++ {
	//	if v, ok := m[nums[i]-1]; ok && v < 2 {
	//		m[nums[i]-1]++
	//	} else if _, ok := m[nums[i]+1]; ok {
	//		m[nums[i]] = m[nums[i]+1] + 1
	//		delete(m, nums[i]+1)
	//	} else if _, ok := m[nums[i]]; !ok {
	//		m[nums[i]] = 1
	//	}
	//}

	//for k, v := range m {
	//
	//}
	//return 0
//}