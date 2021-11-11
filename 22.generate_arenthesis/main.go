package main

import (
	"strings"
)

// 2 <= n <= 8
func generateParenthesis(n int) []string {

	return nil
}

func recursion(n, left, right int, b *strings.Builder, s string) {
	b.WriteString(s)

	if left + right == n {

		return
	}

	if left < n/2 {
		recursion(n, left + 1, right, b, "(")
	}
}