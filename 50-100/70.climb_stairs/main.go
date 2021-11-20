package main

// 对于第n阶，只能通过n-1和n-2通过唯一的方法爬上来
// 所以res[n]=res[n-1]+res[n-2]
func climbStairs(n int) int {
	if n <= 2 { return n }

	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1]+res[i-2]
	}

	return res[n]
}
