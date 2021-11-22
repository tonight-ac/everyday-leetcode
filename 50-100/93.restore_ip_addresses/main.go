package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

var res []string
func restoreIpAddresses(s string) []string {
	res = make([]string, 0)

	recursion(s, make([]string, 0))

	return res
}

func recursion(s string, list []string) {
	if s == "" {
		if len(list) == 4 {
			res = append(res, strings.Join(list, "."))
		}
		return
	}

	for i := 1; i <= 3 && i <= len(s); i++ {
		if isValidField(s[0:i]) {
			recursion(s[i:], append(list, s[0:i]))
		}
	}
}

func isValidField(s string) bool {
	// 去除前导0
	if s[0] == '0' && len(s) != 1 { return false }

	num := 0
	for i := 0; i < len(s); i++ {
		num = num * 10 + int(s[i]-'0')
	}

	return num <= 255
}