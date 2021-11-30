package main

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	count := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if idx, ok := m[nums[i]]; ok {
			delete(m, nums[i])
			if nums[i] <= count[idx][0] {
				count[idx][0] = nums[i]
				m[nums[i]-1] = idx
			} else {
				count[idx][1] = nums[i]
				m[nums[i]+1] = idx
			}
		} else {
			count = append(count, []int{nums[i], nums[i]})
			m[nums[i]-1] = len(count)-1
			m[nums[i]+1] = len(count)-1
		}
	}

	res := 1
	for i := 0; i < len(count); i++ {
		if res < count[i][1] - count[i][0] + 1 {
			res = count[i][1] - count[i][0] + 1
		}
	}

	return res
}

//func longestConsecutive(nums []int) int {
	//m := make(map[int]int)
	//for i := 0; i < len(nums); i++ {
	//	if v, ok := m[nums[i]-1]; ok && v < 2 {
	//		m[nums[i]-1]++
	//	} else if _, ok := m[nums[i]+1]; ok {
	//		m[nums[i]] = m[nums[i]+1] + 1
	//		delete(m, nums[i]+1)
	//	} else if _, ok := m[nums[i]]; !ok {
	//		m[nums[i]] = 1
	//	}
	//}

	//for k, v := range m {
	//
	//}
	//return 0
//}