package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
}

const INT32MAX = 2147483647
const INT32MIN = -2147483648

func divide(dividend int, divisor int) int {
	// 本题中，如果除法结果溢出，则返回 231 − 1。
	// 只有以下一种可能会出现这种结果，直接排除掉
	if dividend == INT32MIN && divisor == -1 {
		return INT32MAX
	}

	// 首先排除符号的干扰，将问题转化为纯粹的减法
	char := 1
	if dividend < 0 {
		char = -char
		dividend = -dividend
	}
	if divisor < 0 {
		char = -char
		divisor = -divisor
	}

	// 符号处理完后，接下来处理商的值
	count := 0
	// 经过测试有个超时的case如下，所以需要考虑性能
	//-2147483648
	//1

	// 不能使用*，/，mod 为了能倍速处理，使用Divisor自身加速
	fastDivisor, fastLevel := divisor, 1
	// 判断条件是否还存在整除的空间
	for	dividend - divisor >= 0 {
		// 如果当前的快速Divisor能够除下，则直接剪掉，并继续增大
		if dividend - fastDivisor >= 0 {
			dividend -= fastDivisor
			count += fastLevel
			fastDivisor, fastLevel = fastDivisor+divisor, fastLevel+1
		} else {
			// 当前快速Divisor已经太大了，直接归0，重新开始
			fastDivisor, fastLevel = divisor, 1
		}
	}

	return char*count
}
