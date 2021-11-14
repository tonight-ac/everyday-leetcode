package main

import "fmt"

func main() {
	list := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(list, 2))
	fmt.Println(list)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 { return 0 }

	// i为待替换的，j为待填充的
	var i, j int
	for i, j = 0, len(nums)-1; i < j; {
		fmt.Println(i, j)
		if nums[i] != val {
			i++
		}
		if nums[j] == val {
			j--
		}
		if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			i++
			j--
		}
	}

	//if nums[i] == val { return 0 }

	return i+1
}
