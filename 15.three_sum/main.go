package main

import (
	"fmt"
	"sort"
)

//输入：
//[-1,0,1,2,-1,-4,-2,-3,3,0,4]
//输出：
//[[-4,0,4],[-4,1,3],[-3,0,3],[-3,1,2],[-2,0,2],[-1,0,1]]
//预期结果：
//[[-4,0,4],[-4,1,3],[-3,-1,4],[-3,0,3],[-3,1,2],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]

//i++[[-4 0 4] [-3 -1 4] [-3 0 3] [-2 -1 3] [-2 0 2] [-1 -1 2] [-1 0 1]]
//j--[[-4 0 4] [-4 1 3] [-3 0 3] [-3 1 2] [-2 0 2] [-1 0 1]]

// -3 4
// -4 3

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	// 先做排序，不排序的话，无法进行，且无法去重
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > 0 { break }

		// 去重 因为i所遍历的后续内容，必定是i-1所遍历的后续内容的子集
		if i > 0 && nums[i] == nums[i-1] { continue }

		// 锁定一个，找到包含它的所有tuple
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i]+nums[l]+nums[r]
			if sum == 0 { // 已经match
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] { l++ } // 目前已经match了，再计算也是重复的，移动到最后一个重复元素
				for l < r && nums[r] == nums[r-1] { r-- } // 目前已经match了，再计算也是重复的，移动到最后一个重复元素
				l++ // 此时再次循环还是重复的，还可以发生match，所以再移动一个
				r-- // 此时再次循环还是重复的，还可以发生match，所以再移动一个
				// 如果都未发生重复，也本身就需要移动，而且由于有序性，不可能同样的数字还有多解法，两个都移动
			} else if sum < 0 { // 三数之和为负，说明较小，需要向大的方向移动
				l++
			} else { // 三数之和为正，说明较大，需要向小的方向移动
				r--
			}
		}
	}

	return res
}

// 二分法失败案例
//func threeSum(nums []int) [][]int {
//	res := make([][]int, 0)
//
//	if len(nums) < 3 {
//		return res
//	}
//
//	//输入：nums = [-1,0,1,2,-1,-4]
//	//输出：[[-1,-1,2],[-1,0,1]]
//
//	// -4 -1 -1 0 1 2
//
//	// 先做排序，不排序的话，无法进行，且无法去重
//	sort.Ints(nums)
//
//	// 首尾两个指针，同时i < j+1需要留出三个数字的空间
//	for i, j := 0, len(nums) - 1; i + 1 < j; {
//		// 求二数之和取反，然后在i+1, j-1的范围内二分找sum
//		sum := -(nums[i] + nums[j])
//
//		// 二分法
//		for l, r := i + 1, j - 1; l <= r; {
//			mid := (l + r) / 2
//			if sum < nums[mid] {
//				r = mid - 1
//			} else if sum > nums[mid] {
//				l = mid + 1
//			} else {
//				// 确实找到了，加入结果集合
//				res = append(res, []int{nums[i], sum, nums[j]})
//				break
//			}
//		}
//
//		// 去除重复元素
//		for i + 1 < j && nums[i] == nums[i+1] {
//			i++
//		}
//		// 去除重复元素
//		for i + 1 < j && nums[j] == nums[j-1] {
//			j--
//		}
//
//		if nums[i] + nums[j] < 0 {
//			// 两数之和为负，说明较小，需要像大的方向移动
//			i++
//		} else if nums[i] + nums[j] > 0 {
//			// 两数之和为正，说明nums[j]较大，需要像小的方向移动
//			j--
//		} else {
//			// 两数相同，不好说先移动哪个，两个都需要尝试
//			// 理论上现在破局只有一招，两个都进去试一下再说
//			// 其实最大的问题在于这里面有0 如何消除0呢？
//			// 其实无法消除，0是临界值，真要命 TODO
//		}
//	}
//
//	return res
//}
