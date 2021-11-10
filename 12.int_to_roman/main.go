package main

import (
	"fmt"
	"strings"
)

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func main() {
	fmt.Println(intToRoman(58))
}

var m = map[int]string{
	1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
}

// 对于 I II III IV V VI VII VIII IX X
func intToRoman(num int) string {
	// num最大3999，所以最多四位
	res := make([]string, 4)

	//输入: num = 1994
	//输出: "MCMXCIV"
	//解释: M = 1000, CM = 900, XC = 90, IV = 4.

	for i, j := 1, 1; num != 0; i, j = i*10, j+1 {
		// 求余数
		mod := num % 10

		// 求罗马数字
		roman := ""
		switch mod {
		case 1,2,3:
			roman = strings.Repeat(m[i], mod)
		case 4:
			roman = m[i]+m[i*5]
		case 5:
			roman = m[i*5]
		case 6,7,8:
			roman = m[i*5]+strings.Repeat(m[i], mod-5)
		case 9:
			roman = m[i]+m[i*10]
		}

		// 保存罗马数字（字符串）
		res[len(res)-j] = roman

		// 处理下一位
		num /= 10
	}

	return strings.Join(res, "")
}
