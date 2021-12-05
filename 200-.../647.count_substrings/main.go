package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
}
func countSubstrings(s string) int {
	res := 0

	// 单中心
	for i := 0; i < len(s); i++ {
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			res++
		}
	}

	// 双中心
	for i := 1; i < len(s); i++ {
		for l, r := i-1, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			res++
		}
	}

	return res
}
