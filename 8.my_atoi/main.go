package main

import (
	"fmt"
	"unicode"
)

const INT32MAX = ^uint32(0)
const INT32MIN = uint32(0)

func main() {
	fmt.Println(int32(INT32MAX))
	fmt.Println(int32(INT32MIN))
}

func myAtoi(s string) int {
	res := int32(0)
	char := int32(1)
	for _, v := range s {
		if unicode.IsDigit(v) {
			if temp := res * 10 + char*(v-'0'); res != temp/10 {
				if char == 1 {
					return int(INT32MAX)
				}
				return int(INT32MIN)
			}
			res *= 10
			res += char*(v-'0')
		} else if unicode.IsSpace(v) {
			continue
		} else if string(v) == "-" {
			char = -1
		} else if string(v) == "+" {
			continue
		} else {
			break
		}
	}

	return int(res)
}
