package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"a","flow","flight"}))
}

//输入：strs = ["flower","flow","flight"]
//输出："fl"

// 使用一个退化的字典树
func longestCommonPrefix(strs []string) string {
	// 选取第一个作为模版，不必判空，入参必有一个
	p := strs[0]

	// 计算最大长度
	length := len(p)

	for i := 0; i < len(strs); i++ {
		s := strs[i]

		// 存在空串就不需要比较了，因为没有前缀
		if len(s) == 0 {
			return ""
		}

		j := 0
		for ; j < length && j < len(s); j++ {
			if s[j] != p[j] {
				break
			}
		}

		// 循环结束时计算最大前缀
		if j < length {
			length = j
		}
	}

	return p[0:length]
}