package main

func lengthOfLongestSubstring(s string) int {
	// 只可能是英文字母、数字、符号、空格
	// map可以动态扩容，大部分字符串应该比较短，我们先开15
	m := make(map[int32]int, 15)

	// 最大值，计数器
	max, count := 0, 0

	for i, v := range s {
		// 目前遍历到的元素曾经出现过，计算一下是否是出现在当前range的范围
		if idx, ok := m[v]; ok && i - idx <= count {
			count = i - idx // 是 缩小范围重新计数
		} else {
			count++ // 不是 加一
		}

		m[v] = i // 更新最近出现的下标map

		// 更新最大值
		if count > max {
			max = count
		}
	}

	return max
}
