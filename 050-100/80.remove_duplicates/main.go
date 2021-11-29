package main

import "fmt"

func main()  {
	nums := []int{0,0,1,1,1,1,2,3,3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
//输入：nums = [1,1,1,2,2,3]
//输出：5, nums = [1,1,2,2,3]
//
func removeDuplicates(nums []int) int {
	// 本来就有序长度又小于2，直接返回
	if len(nums) == 1 { return 1 }
	res := 2
	for i := 2; i < len(nums); i++ {
		// i表示待插入元素，res表示插入位置，i不能和上一个和上两个插入的相同
		// 简化一下，由于有序性，nums[res-2]必然小于nums[res-1]，所以只判断nums[res-2]足够
		if nums[i] != nums[res-2] {
			nums[res] = nums[i]
			res++
		}
	}

	return res
}

