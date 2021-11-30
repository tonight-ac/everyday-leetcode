package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	// 清洗s
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if unicode.IsDigit(r) || unicode.IsLower(r) {
			b.WriteRune(r)
		} else if unicode.IsUpper(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}

	// 判断是否回文
	ss := b.String()
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		if ss[i] != ss[j] {
			return false
		}
	}

	return true
}
