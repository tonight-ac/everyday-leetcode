package main

//输入：
//[1,1,2]
//输出：
//[2,1]
//预期结果：
//[1,2]

func removeDuplicates(nums []int) int {
	// 长度0，直接返回
	if len(nums) == 0 { return 0 }

	// 结果数组的末尾
	end := 0

	for i := 0; i < len(nums); i++ {
		// 不相同就往里添加
		if nums[i] != nums[end] {
			end++
			nums[end] = nums[i]
		}
	}

	return end+1
}
