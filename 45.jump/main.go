package main

import "fmt"

func main()  {
	fmt.Println(jump([]int{4,1,1,3,1,1,1}))
}

// 另一种方案，单纯下一个节点跳得远没用，还需要考虑到当前移动的节点数目
func jump(nums []int) int {
	// 如果只有一个元素，不用跳了
	if len(nums) == 1 { return 0 }
	count := 0 // 每循环一轮count会+1，表示换了一个新的节点i，进行了一次跳跃
	for i := 0; i < len(nums); count++ {
		// 下一跳就可以成功了
		if i + nums[i] >= len(nums)-1 {
			// 因为最后一次没有真跳，所以加上最后一次
			return count + 1
		}
		// 开始在接下来的范围内进行搜索
		max, idx := nums[i+1] + 1, i + 1
		for j := i + 2; j <= i + nums[i]; j++ {
			// 当前移动节点数目+未来可移动范围排大小
			if max <= (j - i) + nums[j] {
				max, idx = (j - i) + nums[j], j
			}
		}

		// 跳到下一个位置
		i = idx
	}

	return count
}

// 每次都选择可跳跃范围内最大的地方跳
// 这样的话也能保证每次的可搜索范围是最大的，相辅相成保证跳的次数更少

// 这种方案失败，下面的case无法通过
//输入: nums = [10,9,8,7,6,5,4,3,2,1,1,0]
//输出: 2
//func jump(nums []int) int {
//	// 如果只有一个元素，不用跳了
//	if len(nums) == 1 { return 0 }
//	count := 0
//	for i := 0; i < len(nums); count++ {
//		if i + nums[i] >= len(nums)-1 {
//			// 因为最后一次没有真跳，所以加上最后一次
//			return count + 1
//		}
//		// 开始在接下来的范围内进行搜索
//		max, idx := nums[i+1], i + 1
//		for j := i + 2; j <= i + nums[i]; j++ {
//			// 一定是小于等于，小于那必然更新，如果等于的话必然选更靠后的
//			if max <= nums[j] {
//				max, idx = nums[j], j
//			}
//		}
//		i = idx
//	}
//
//	return count
//}
