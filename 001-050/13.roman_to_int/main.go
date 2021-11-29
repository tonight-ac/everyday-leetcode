package main

var m = map[uint8]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	max := 0

	// 根据罗马数字规范：
	// 如果一个小的数字在大的数字左侧，必定是做减法，其他情况都需要做加法
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		flag := -1 // 默认需要减法
		if cur >= max {
			// 记录当前最大值
			max = cur
			flag = 1 // 改为加法
		}
		res += flag*cur
	}

	return res
}
