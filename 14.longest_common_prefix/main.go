package main

//输入：strs = ["flower","flow","flight"]
//输出："fl"

func longestCommonPrefix(strs []string) string {
	// 既然len(strs) > 0, 怎选取第一个作为模版
	p := strs[0]

	// 对p的每个字符计数用的list
	counts := make([]int, len(p))

	// 计算最大长度
	max := 0

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		// 存在空串就不需要比较了
		if len(s) == 0 { return "" }
		for j := 0; j < len(s) && j < len(p); j++ {
			if s[j] == p[j] {
				counts[j]++
			}
		}
	}

	return p[0:max]
}
