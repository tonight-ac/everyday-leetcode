package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("aa", "aa"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) == 1 { return false }

	unMatch := make([]int, 0)

	uni := make(map[byte]interface{})
	for i := 0; i < len(s); i++ {
		uni[s[i]-'a'] = nil
		if s[i] != goal[i] {
			unMatch = append(unMatch, i)
		}
	}

	if len(unMatch) == 2 { // 不匹配数量相等，看看交换能不能解决问题
		return s[unMatch[0]] == goal[unMatch[1]] && s[unMatch[1]] == goal[unMatch[0]]
	} else if len(unMatch) > 2 || len(unMatch) == 1 { // 大于2或者等于，则靠一次交换不能解决问题
		return false
	}

	// 如果为真，s必然存在重复元素，重复元素原地交换即可
	return len(uni) < len(s)
}
