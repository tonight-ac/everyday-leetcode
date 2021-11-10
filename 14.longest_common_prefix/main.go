package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"a","flow","flight"}))
}

//输入：strs = ["flower","flow","flight"]
//输出："fl"

func longestCommonPrefix(strs []string) string {
	// 选取第一个作为模版，不要判空，入参必有一个
	p := strs[0]

	// 对p的每个字符计数用的list
	// int8节约内存，题目规定入参数组长度不超过200
	counts := make([]int8, len(p))

	// 计算最大长度
	min := len(p)

	for i := 0; i < len(strs); i++ {
		s := strs[i]

		// 存在空串就不需要比较了，因为没有前缀
		if len(s) == 0 { return "" }

		j := 0
		for ; j < len(s) && j < len(p); j++ {
			if s[j] != p[j] {
				break
			}
			counts[j]++
		}

		// 循环结束时计算最大前缀
		if j < min {
			min = j
		}
	}

	return p[0:min]
}