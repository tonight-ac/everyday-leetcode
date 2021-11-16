package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply("123", "456"))
}

// 111
// 10
func multiply(num1 string, num2 string) string {
	res := "0"

	// 把num1表示二者短的，num2表示二者长的
	// 以此来确定尽量外层循环的次数更少一些
	if len(num1) > len(num2) {
		temp := num1
		num1 = num2
		num2 = temp
	}

	for i := len(num1)-1; i >= 0; i-- {
		// 0直接跳过
		if num1[i] == '0' { continue }
		for j := len(num2)-1; j >= 0; j-- {
			// 0直接跳过
			if num1[j] == '0' { continue }

			n := (num1[i]-'0') * (num2[j]-'0')
			// 二者相乘再乘以系数
			var temp strings.Builder
			if n/10 != 0 {
				temp.WriteByte(n/10+'0')
			}
			if n%10 != 0 {
				temp.WriteByte(n%10+'0')
				temp.WriteString(strings.Repeat("0", len(num1)-1-i))
			}

			// 做一个字符串加法
			res = add(res, temp.String())
		}
	}

	return res
}

func add(n1, n2 string) string {
	res := ""

	for i, j, carry := len(n1)-1, len(n2)-1, byte(0); i >= 0 || j >= 0 || carry != 0; i, j = i-1,j-1 {
		v1, v2 := byte(0), byte(0)
		if i >= 0 { v1 = n1[i] - '0' }
		if j >= 0 { v2 = n2[j] - '0' }
		tot := v1 + v2 + carry

		res = string(tot % 10 + '0') + res
		carry = tot/10
	}

	return res
}