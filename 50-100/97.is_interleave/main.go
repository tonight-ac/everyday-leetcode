package main

import "fmt"

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbcbbcac"))
}

// 大力没有出奇迹，而是超时了
//"aabcc"
//"dbbca"
//"aadbcbbcac"
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) { return false }

	return recursion(s1, s2, s3, 0, 0)
}

func recursion(s1 string, s2 string, s3 string, left, right int) bool {
	if len(s1) + len(s2) + len(s3) == 0 {
		return abs(left-right) <= 1
	}

	for i := 0; i < len(s1) && i < len(s3) && s1[i] == s3[i]; i++ {
		if ok := recursion(s1[i+1:], s2, s3[i+1:], left+1, right); ok {
			return true
		}
	}

	for i := 0; i < len(s2) && i < len(s3) && s2[i] == s3[i]; i++ {
		if ok := recursion(s1, s2[i+1:], s3[i+1:], left, right+1); ok {
			return true
		}
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}