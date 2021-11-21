package main

import "fmt"

func main() {
	nums1 := []int{2,0}
	merge(nums1, 1, []int{1}, 1)
	fmt.Println(nums1)
}

// 选取两个中更小的放入
// 先填入nums1高位，因为高位是空白的
func merge(nums1 []int, m int, nums2 []int, n int) {
	end := m+n-1
	i, j := m-1, n-1
	for ; i >= 0 && j >= 0; {
		if nums1[i] < nums2[j] {
			nums1[end] = nums2[j]
			j--
		} else {
			nums1[end] = nums1[i]
			i--
		}
		end--
	}
	// 可能存在nums2没有填完而nums1先填完的情况
	if j >= 0 {
		for ; j >= 0; end, j = end-1, j-1 { nums1[end] = nums2[j] }
	}
}
