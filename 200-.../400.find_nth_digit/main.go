package main

import "fmt"

func main() {
	fmt.Println(findNthDigit(1000))
}
// n为位数
// (n-1)*10
// 9 90 900 9000
// 先确定位数
func findNthDigit(n int) int {
	if n < 10 { return n }

	nums := 1
	length := 1
	// 去掉个宽度为length的9*nums个数据
	for n > nums * 9 * length {
		n -= nums * 9 * length
		nums *= 10
		length++
	}

	// 此时必然在pos到pos*10之间
	for n > length {
		n -= length
		nums++
	}

	// 此时只需要提取出nums的高第n位就行
	length -= n // 转化为低第length位
	for {
		if length == 0 {
			return nums%10
		}
		length--
		nums/=10
	}
}

