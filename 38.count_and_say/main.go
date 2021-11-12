package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAndSay(4))
}

// 存储一下结果
// n最大30
var temp [31]string
func countAndSay(n int) string {
	if n == 1 { return "1" }

	sub := temp[n-1]
	if sub == "" {
		sub = countAndSay(n-1)
		temp[n-1] = sub
	}

	// 会出现大量字符串拼接，使用strings.Builder提高性能&节约内存
	var b strings.Builder
	count, cur := byte(1), sub[0]

	for i := 1; i < len(sub); i++ {
		if sub[i] == cur {
			count++
		} else {
			b.WriteByte(count+'0')
			b.WriteByte(cur)
			count, cur = byte(1), sub[i]
		}
	}

	b.WriteByte(count+'0')
	b.WriteByte(cur)

	return b.String()
}
