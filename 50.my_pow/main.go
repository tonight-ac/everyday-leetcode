package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, -2))
}

// 其实就是实现一个快速幂运算
// 原理：x^22可以改写成x^16*x^4*x^2
func myPow(x float64, n int) float64 {
	if n == 0 { return 1.0 }

	res := 1.0
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}
	for n > 0 {
		// 最低位存在的话，乘一个x
		if n & 1 == 1 { res = res * x }
		// 对x进行翻倍
		x = x * x
		// 吃掉最低位
		n >>= 1
	}

	if flag {
		return 1.0 / res
	}
	return res
}

// 简单的慢速方法
//func myPow(x float64, n int) float64 {
//	if n == 0 { return 1.0 }
//	res := x
//	if n > 0 {
//		for i := 0; i < n-1; i++ {
//			res *= x
//		}
//	} else {
//		for i := 0; i < -(n-1); i++ {
//			res /= x
//		}
//	}
//
//	return res
//}
