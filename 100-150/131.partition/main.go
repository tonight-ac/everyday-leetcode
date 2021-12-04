package main

// #a#a#b#
// #b#a#a#
// a a b aa


// 其实主要问题在于怎么切分
// 然后对切分的结果做判断
func partition(s string) [][]string {

	return nil
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] { return false }
	}

	return true
}