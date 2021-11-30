package main

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]-1]; ok && v < 2 {
			m[nums[i]-1]++
		} else if _, ok := m[nums[i]+1]; ok {
			m[nums[i]] = m[nums[i]+1] + 1
			delete(m, nums[i]+1)
		} else if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		}
	}

	//for k, v := range m {
	//
	//}
	return 0
}