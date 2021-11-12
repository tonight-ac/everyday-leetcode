package main

const INT32MAX = 2147483647
const INT32MIN = -2147483648

func divide(dividend int, divisor int) int {
	// 本题中，如果除法结果溢出，则返回 231 − 1。
	if dividend == INT32MIN && divisor == -1 {
		return INT32MAX
	}

	// timeout
	//-2147483648
	//1

	// 处理符号
	char := 1
	if dividend < 0 {
		char *= -1
	}
	if divisor < 0 {
		char *= -1
	}

	// 处理商
	count := 0
	dividend, divisor = abs(dividend), abs(divisor)
	for	dividend - divisor >= 0 {
		count ++
		dividend -= divisor
	}

	return char*count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}