package main

import "fmt"

// rabbbit
// rabbit
// 3
//"adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
//"bcddceeeebecbc"
func main() {
	//fmt.Println(numDistinct("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t) + 1)
	for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(s) + 1) }
	for i := 0; i < len(dp[0]); i++ { dp[0][i] = 1 }

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if t[i - 1] == s[j - 1] {
				// 选与不选
				dp[i][j] = dp[i - 1][j - 1] + dp[i][j - 1]
			}  else {
				// 没有权利，只能不选
				dp[i][j] = dp[i][j - 1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
//   a b c d
// a 1 1 1
// b
// d

//   e e e
// e 1 2 3
// e 0 1 3
// e 0 0 1
// 如果当前s不要还能匹配说明是一种新情况
//   r a b b b i t
// r 1 1 1 1 1 1 1
// a 0 1 1 1 1 1 1
// b 0 0 1 2 3 3 3
// b 0 0 0 1 3 3 3
// i 0 0 0 0 0 1 1
// t 0 0 0 0 0 0 1
// dp[i][j-1]为0 dp[i][j] = 1
// dp[i][j-1]不为0
//		还需要区分情况
//			t[i] != t[i-1] dp[i][j] += dp[i][j-1]
// 			t[i] == t[i-1] dp[i][j] += dp[i][j-1] dp[i][j] += dp[i-1][j-1]
//func numDistinct(s string, t string) int {
//	// 初始化dp
//	dp := make([][]int, len(t))
//	for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(s)) }
//
//	for i := 0; i < len(s); i++ {
//		if t[0] == s[i] { dp[0][i] = 1 }
//		if i > 0 { dp[0][i] += dp[0][i-1] }
//	}
//
//	for i := 1; i < len(t); i++ {
//		for j := i; j < len(s); j++ {
//			if t[i] == s[j] {
//				if dp[i][j-1] == 0 {
//					// 初次匹配
//					dp[i][j] = 1
//				} else {
//					dp[i][j] += dp[i][j-1]
//					if t[i] == t[i-1] {
//						dp[i][j] += dp[i-1][j-1]
//					}
//				}
//			} else {
//				dp[i][j] = dp[i][j-1]
//			}
//		}
//	}
//
//	res := 1
//	for i := 0; i < len(t); i++ {
//		min := dp[i][len(s)-1]
//		for i < len(t) - 1 && t[i] == t[i+1] {
//			i++
//			if min > dp[i][len(s)-1] {
//				min = dp[i][len(s)-1]
//			}
//		}
//		res *= min
//	}
//
//	return res
//}

// 搜索超时
//var counts [][]int
//var res int
//func numDistinct(s string, t string) int {
//	res = 0
//	counts = make([][]int, 0)
//	for i := 0; i < len(t); i++ {
//		//if i > 0 && t[i] == t[i-1] {
//		//	continue
//		//}
//		count := make([]int, 0)
//		for j := 0; j < len(s); j++ {
//			// 连续相等能不能压缩
//			if t[i] == s[j] {
//				count = append(count, j)
//			}
//		}
//		counts = append(counts, count)
//	}
//
//	recursion(0, -1)
//
//	return res
//}
//
//func recursion(idx, prev int) {
//	if idx == len(counts) {
//		res++
//		return
//	}
//	// 二分搜索起始下标
//	l, r := 0, len(counts[idx]) - 1
//	mid := 0
//	for l <= r {
//		mid = (l+r)/2
//		if counts[idx][mid] > prev {
//			r = mid-1
//		} else if counts[idx][mid] < prev {
//			l = mid+1
//		} else {
//			l = mid+1
//			break
//		}
//	}
//	for i := l; i < len(counts[idx]); i++ {
//		if counts[idx][i] > prev {
//			recursion(idx+1, counts[idx][i])
//		}
//	}
//}

//d不出来
//func numDistinct(s string, t string) int {
//	// 前i(in t)中包含前j(in s)的多少位前缀
//	dp := make([][]int, len(t))
//	for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(s)) }
//
//	for i := 0; i < len(dp); i++ {
//		for j := 0; j < len(dp[i]); j++ {
//			if i == 0 {
//				if t[i] == s[j] {
//					dp[i][j] = 1
//				}
//				if j > 0 {
//					dp[i][j] += dp[i][j-1]
//				}
//			} else if t[i] == s[j] {
//				dp[i][j] += dp[i-1][j]
//			} else if j > 0 {
//				dp[i][j] += dp[i][j-1]
//			}
//		}
//	}
//
//	return dp[len(t)-1][len(s)-1]
//}
