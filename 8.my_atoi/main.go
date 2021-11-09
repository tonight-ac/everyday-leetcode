package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(INT32MAX)
	fmt.Println(INT32MIN)
}

const INT32MAX = 2147483647
const INT32MIN = -2147483648

func myAtoi(s string) int {
	// 处理前缀

	// 处理数字

	res := int32(0)
	char, setChar := int32(1), false
	for _, v := range s {
		if unicode.IsDigit(v) {
			temp := res*10+char*(v-'0')
			if res != temp/10 {
				if char == 1 {
					return INT32MAX
				}
				return INT32MIN
			}
			res = temp
		} else if unicode.IsSpace(v) {
			continue
		} else if string(v) == "-" {
			if setChar { return 0 }
			char, setChar = -1, true
		} else if string(v) == "+" {
			if setChar { return 0 }
			char, setChar = 1, true
			continue
		} else {
			break
		}
	}

	return int(res)
}
