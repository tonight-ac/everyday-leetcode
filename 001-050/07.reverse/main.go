package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(-123))
}

// int32 范围 -2147483648 ~ 2147483647
// 2147483647
// 1111111119
// list: 9111111111
// res: 911111111
// 9111111110 + 1
// 911111111 -> 9111111111
// 911111111 * 10 + 1 => -xxxxxx
// -xxxxxx / 10 = -xxxxx
// 123 * 10 + 4 = 1234
// 1234 / 10 = 123
func reverse(x int) int {
	if x > -10 && x < 10 { return x }

	res := int32(0)
	list := make([]int32, 0)

	// 对x进行拆分，将x每一位放入list中
	// x=123
	// list=[3,2,1]
	// x=-123
	// list[-3,-2,-1]
	for ; x != 0; x/=10 {
		list = append(list, int32(x%10))
	}

	// 将list中的每一位，组装回res中
	for i := 0; i < len(list); i++ {
		// 溢出操作不可逆，所以我们做逆操作，避免溢出
		if temp := res * 10 + list[i]; res != temp/10 {
			return 0
		}
		res = res * 10 + list[i]
	}

	return int(res)
}
