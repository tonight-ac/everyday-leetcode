package main

import "fmt"

// 和45题类似，但是45是保证可达的情况下贪心计算最少次数，这次是对可达性的探索
// 逻辑相同把45题拿过来改一下就行
//输入：
//[3,2,1,0,4]
//输出：
//true
//预期结果：
//false
func main() {
	fmt.Println(canJump([]int{3,2,1,0,4}))
}

func canJump(nums []int) bool {
	// 如果只有一个元素，不用跳了
	if len(nums) == 1 { return true }
	for i := 0; i < len(nums); {
		// 当前可跳跃次数为0，并且不是最后一个节点，返回false
		if i != len(nums)-1 && nums[i] == 0 {
			return false
		}
		// 下一跳就可以成功了
		if i + nums[i] >= len(nums)-1 {
			return true
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

	return false
}
