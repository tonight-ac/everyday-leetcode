package main

import "fmt"

// 0,1,2,2,2,2,2,3,4
// 0 1 4 3
func main() {
	list := []int{0,1,2,2,2,2,2,3,4}
	fmt.Println(removeElement(list, 2))
	fmt.Println(list)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 { return 0 }

	// i为待替换的，j为待填充的
	for i, j := 0, len(nums)-1; i < j; {
		// 可以换位置，更换
		if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			nums[j] = val
			i++
			j--
		}
		// 不符合移动i向后寻找待更换位置
		if nums[i] != val {
			i++
		}
		if nums[j] == val {
			j--
		}
	}

	// 通过i、j去计算结果，不方便，干脆单独计算
	res := 0
	for _, v := range nums {
		if v != val {
			res++
		} else {
			break
		}
	}

	return res
}