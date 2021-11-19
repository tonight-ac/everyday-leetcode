package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	if x == 0 { return 0 }

	res := int32(0)
	list := make([]int32, 0)

	for ; x != 0; x/=10 {
		list = append(list, int32(x%10))
	}

	for i := 0; i < len(list); i++ {
		// 溢出操作不可逆，所以我们做逆操作，避免溢出
		if temp := res * 10 + list[i]; res != temp/10 {
			return 0
		}
		res = res * 10 + list[i]
	}

	return int(res)
}
