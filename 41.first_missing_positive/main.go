package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))
}

//输入：nums = [1,2,0]
//输出：3
func firstMissingPositive(nums []int) int {
	//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
	//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

	// 找到最小的连续正整数
	length := len(nums)

	// 核心技巧：nums内的数字最多为纯正数，并且刚好为从1到len(nums)，这样结果就是len(nums)+1
	// 除此之外，唯一可能的解数值必然小于len(nums)
	for i := 0; i < length; i++ {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i] - 1] != nums[i] {
			// 满足在指定范围内、并且没有放在正确的位置上，才交换
			// 例如：数值 3 应该放在索引 2 的位置上
			temp := nums[nums[i] - 1]
			nums[nums[i] - 1] = nums[i]
			nums[i] = temp
		}
	}

	// [1, -1, 3, 4]
	for i := 0; i < length; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}

	// 都正确则返回数组长度 + 1
	return length + 1
}
