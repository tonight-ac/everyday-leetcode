package main

import "strings"

func toLowerCase(s string) string {
	var b strings.Builder
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c = c-'A'+'a'
		}
		_, _ = b.WriteRune(c)
	}
	return b.String()
}


