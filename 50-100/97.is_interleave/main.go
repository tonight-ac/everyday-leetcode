package main

import (
	"fmt"
	"strings"
)
//"abbbbbbcabbacaacccababaabcccabcacbcaabbbacccaaaaaababbbacbb"
//"ccaacabbacaccacababbbbabbcacccacccccaabaababacbbacabbbbabc"
//"cacbabbacbbbabcbaacbbaccacaacaacccabababbbababcccbabcabbaccabcccacccaabbcbcaccccaaaaabaaaaababbbbacbbabacbbacabbbbabc"
//"cacbbbaaabbacbbbbabbcaccccabaaacacbcaacababbaabbaccacbaabac"
//"cbcccabbbbaaacbaccbabaabbccbbbabacbaacccbbcaabaabbbcbcbab"
//"ccbcccacbabbbbbbaaaaabbaaccbabbbbacbcbcbaacccbacabbaccbaaabcacbbcabaabacbbcaacaccbbacaabababaabbbaccbbcacbbacabbaacb"
//"aabcc"
//"dbbca"
//"aadbbbaccc"
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true
func main() {
	fmt.Println(isInterleave("cacbbbaaabbacbbbbabbcaccccabaaacacbcaacababbaabbaccacbaabac", "cbcccabbbbaaacbaccbabaabbccbbbabacbaacccbbcaabaabbbcbcbab", "ccbcccacbabbbbbbaaaaabbaaccbabbbbacbcbcbaacccbacabbaccbaaabcacbbcabaabacbbcaacaccbbacaabababaabbbaccbbcacbbacabbaacb"))
	//fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

var temp []string
var ss string
// 先对于更短的串和混合串进行匹配拆解，将空出来的所有可能拼接然后和长串进行比较
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) { return false }

	// s1作为更短的串，s2作为更长的串
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	temp = make([]string, 0)

	ss = s2

	recursion(s1, s3, make([]string, 0))

	//fmt.Println(temp)

	// 将所有可能和s2进行比较
	for _, v := range temp {
		if v == s2 {
			return true
		}
	}

	return false
}

func recursion(s1, s3 string, s []string) {
	// s1匹配完毕了，把结果保存后返回
	if len(s1) == 0 {
		temp = append(temp, strings.Join(s, "")+s3)
		return
	}

	// s3也为空了，返回，不然会越界
	if len(s3) == 0 {
		return
	}

	// 关键性剪枝操作
	if !strings.HasPrefix(ss, strings.Join(s, "")) {
		return
	}

	if s1[0] == s3[0] {
		// 相等并匹配
		recursion(s1[1:], s3[1:], s)
		// 相等但不让s1匹配，跳过
		recursion(s1, s3[1:], append(s, s3[0:1]))
	} else {
		// 不相等，直接跳过
		recursion(s1, s3[1:], append(s, s3[0:1]))
	}
}

// 改dp试一下
//func isInterleave(s1 string, s2 string, s3 string) bool {
//	dp := make([][]int, len(s1))
//	for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(s2)) }
//
//
//
//	return false
//}
//
//func abs(a int) int {
//	if a < 0 {
//		return -a
//	}
//	return a
//}

// 大力没有出奇迹，而是超时了
//"aabcc"
//"dbbca"
//"aadbcbbcac"
//func isInterleave(s1 string, s2 string, s3 string) bool {
//	if len(s1) + len(s2) != len(s3) { return false }
//
//	return recursion(s1, s2, s3, 0, 0)
//}
//
//func recursion(s1 string, s2 string, s3 string, left, right int) bool {
//	if len(s1) + len(s2) + len(s3) == 0 {
//		return abs(left-right) <= 1
//	}
//
//	if left <= right+1 {
//		for i := 0; i < len(s1) && i < len(s3) && s1[i] == s3[i]; i++ {
//			if ok := recursion(s1[i+1:], s2, s3[i+1:], left+1, right); ok {
//				return true
//			}
//		}
//	}
//
//	if right <= left+1 {
//		for i := 0; i < len(s2) && i < len(s3) && s2[i] == s3[i]; i++ {
//			if ok := recursion(s1, s2[i+1:], s3[i+1:], left, right+1); ok {
//				return true
//			}
//		}
//	}
//
//	return false
//}
