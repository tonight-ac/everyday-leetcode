package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1,2,3,5,4}
	nextPermutation(list)
	fmt.Println(list)
}

// 参考46题全排列
// 寻找一个数字
// 要求其后面的数字都大于他（大于不是大于等于）
// 同时要求越靠后越好
// 然后将最小的大于他的数字和他交换
// 再把剩余的按照升序排列
func nextPermutation(nums []int) {
	if len(nums) == 1 { return }

	// 从后向前扫描，找到第一个nums[i] > nums[i-1]的所在位置
	tarIdx := -1
	for i := len(nums)-1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			tarIdx = i - 1
			break
		}
	}

	// 临界条件，整体为降序，需要反转数组
	if tarIdx == -1 {
		for i, j := 0, len(nums)-1; i < j; i, j = i + 1, j - 1 {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
		}
		return
	}

	// 挑选从idx之后比nums[i-1]大的最小的元素
	min, minIdx := nums[tarIdx+1], tarIdx+1
	for i := tarIdx + 1; i < len(nums); i++ {
		if nums[i] >= nums[tarIdx] && nums[i] <= min {
			min, minIdx = nums[i], i
		}
	}

	// 交换
	temp := nums[tarIdx]
	nums[tarIdx] = nums[minIdx]
	nums[minIdx] = temp

	// 对tarIdx后面的进行排序
	sort.Ints(nums[tarIdx+1:])

	return
}
