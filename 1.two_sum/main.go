package main

// 必然存在一个答案，且仅存在一个答案
// 数组中同一个元素在答案里不能重复出现

// 计算数值之和 获取下标

// 1 向后搜索 a+b=target

// 2 向前搜索 a=target-b

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

func twoSum(nums []int, target int) []int {
	// 9
	// 2 7 11 15
	// 2 7
	// 2 11
	// 2 15
	// 7 11
	// 7 15
	// 11 15

	// nums[k]+nums[i]=target
	// nums[k]=target-nums[i]
	// nums[k] k
	// map[int]int

	// 新开一个map
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		// 2 7 11 15
		// 2
		// 2->0
		// 2
		// 2 11 15 7
		//
		if k, ok := m[target-nums[i]]; ok {
			return []int{k, i}
		}

		m[nums[i]] = i
	}

	return nil
}

//func twoSum(nums []int, target int) []int {
//	m := make(map[int]int, len(nums))
//
//	for i, v := range nums {
//		if idx, ok := m[target-v]; ok {
//			return []int{idx, i}
//		}
//		m[v] = i
//	}
//
//	return nil
//}
