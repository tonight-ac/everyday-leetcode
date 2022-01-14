package main

import "fmt"

func main() {
	fmt.Println(increasingTriplet([]int{20,100,10,12,5,13}))
}

//输入：
//[20,100,10,12,5,13]
//输出：
//false
//预期结果：
//true

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 { return false }

	var first, second *int
	for i := 0; i < len(nums); i++ {
		if first == nil || nums[i] < *first {
			first = &nums[i]
		} else if nums[i] > *first && (second == nil || nums[i] < *second) {
			second = &nums[i]
		} else if second != nil && nums[i] > *second {
			return true
		}
	}

	return false
}
