package main

// 我理解就是50题的反向操作
func mySqrt(x int) int {
	res := 0
	n := 2
	//count := 1
	for x > 0 {
		// 最低位存在的话，乘一个x
		if x & 1 == 1 {
			x = x / n
		}
		// 对x进行翻倍
		n = n * n
		// 吃掉最低位
		x >>= 1
	}

	return res
}
