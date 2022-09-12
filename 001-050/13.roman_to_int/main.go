package main

func main() {
	romanToInt("MCMXCIV")
}

var m = map[uint8]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}


// IV 5-1=4
// IX 10-1=9
// VI 1+5=6
// VII 1+1+5 = 7
// VIII
func romanToInt(s string) int {
	res := 0
	max := 0

	//输入: s = "MCMXCIV"
	//输出: 1994
	//解释: M = 1000, CM = 900, XC = 90, IV = 4.

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
