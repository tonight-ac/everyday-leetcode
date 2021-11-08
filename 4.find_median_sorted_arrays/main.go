package main

func main() {
	l1 := []int{1,3}
	l2 := []int{2}
	findMedianSortedArrays(l1, l2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tot := len(nums1) + len(nums2)
	list := make([]int, 0)

	// 毕竟是有序的，使用栈的思路
	for i, j := 0, 0; len(list) < tot/2+1; {
		v1, v2 := 9999999, 9999999
		if i < len(nums1) {
			v1 = nums1[i]
		}
		if j < len(nums2) {
			v2 = nums2[j]
		}
		if v1 < v2 {
			list = append(list, v1)
			i++
		} else {
			list = append(list, v2)
			j++
		}
	}

	if tot%2 == 0 {
		// 偶数
		return float64(list[len(list)-1] + list[len(list)-2])/2.0
	} else {
		// 奇数
		return float64(list[len(list)-1])
	}
}
