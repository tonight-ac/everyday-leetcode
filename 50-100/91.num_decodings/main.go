package main

import "fmt"

func main() {
	fmt.Println(numDecodings("1123"))
}

// 换一个dp的方式
func numDecodings(s string) int {
	dp := make([]int, len(s))

	// 含有导0，不合法，返回0
	if s[0] == '0' { return 0 }

	// 合法，置第一个为1
	dp[0] = 1

	for i := 1; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' { // 如果当前合法
			if s[i-1] != '0' && (s[i-1]-'0')*10+(s[i]-'0') <= 26 { // 一位和两位都合法，+2
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else { // 两位不合法
				dp[i] = dp[i-1]
			}
		} else { // 当前不合法
			if s[i-1] != '0' && (s[i-1]-'0')*10+(s[i]-'0') <= 26 { // 和前一位组成的两位合法
				if i >= 2 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 1
				}
			} else { // 一位和两位都不合法
				break
			}
		}
	}

	return dp[len(s)-1]
}

// 递归解决，递归超时了
// case "111111111111111111111111111111111111111111111"
//var res int
//func numDecodings(s string) int {
//	res = 0
//	recursion(s)
//	return res
//}
//
//func recursion(s string) {
//	if len(s) == 0 {
//		res++
//		return
//	}
//
//	if s[0] != '0' {
//		recursion(s[1:])
//	}
//	if len(s) > 1 && s[0] != '0' && (s[0]-'0')*10+(s[1]-'0') <= 26 {
//		recursion(s[2:])
//	}
//}