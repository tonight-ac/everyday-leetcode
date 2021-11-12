package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 2, 8, 8, 9, 10}, 8))
}

var res = make([]int, 2)
// 以二分法为基础做一定的改进 TODO 未上传，且待优化
func searchRange(nums []int, target int) []int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			for nums[l] != target {
				l++
			}
			for nums[r] != target {
				r--
			}
			res[0], res[1] = l, r
			return res
		}
	}
	res[0], res[1] = -1, -1
	return res
}
