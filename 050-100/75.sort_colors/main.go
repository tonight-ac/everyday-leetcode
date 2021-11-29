package main

import "fmt"

func main() {
	nums := []int{0,0,1}
	sortColors(nums)
	fmt.Println(nums)
}

// 寻找i和j下一个可交换位置
func sortColors(nums []int) {
	// 先换0到前面
	for l, i := 0, 1; i < len(nums); i++ {
		if nums[i] == 0 {
			for l < i && nums[l] == 0 { l++ }
			swap(nums, l, i)
		}
	}

	// 再换2到后面
	for r, j := len(nums)-1, len(nums)-2; j >= 0; j-- {
		if nums[j] == 2 {
			for r > j && nums[r] == 2 { r-- }
			swap(nums, r, j)
		} else if nums[j] == 0 {
			break // 因为0都已经归位了
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 这个解法特殊case太多
// 把0往左放，把2往右放，1自然往中间挤
//func sortColors(nums []int) {
//	for i, j := 0, len(nums)-1; ; {
//		if nums[i] == 2 {
//			for i < j && nums[j] == 2 { j-- }
//			swap(nums, i, j)
//			j--
//		} else if nums[i] == 0 {
//			i++
//		}
//		// 因为上一轮刚修改过i和j 所以需要对i和j的合法性做判断
//		if j >= 0 && i < len(nums)-1 {
//			if nums[j] == 0 {
//				for i < j && nums[i] == 0 { i++ }
//				swap(nums, i, j)
//				i++
//			} else if nums[j] == 2 {
//				j--
//			}
//		}
//		if i >= j-1 {
//			break
//		}
//	}
//}
