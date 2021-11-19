package main

func lengthOfLastWord(s string) int {
	// 从后面开始遍历
	i := len(s)-1

	res := 0

	// 排除所有后置空格
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	for ; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		res++
	}

	return res
}
