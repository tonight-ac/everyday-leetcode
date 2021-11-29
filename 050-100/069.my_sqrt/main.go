package main

// 使用二分法试一下提升速度
func mySqrt(x int) int {
	l, r, mid := 0, x, 0

	for l <= r {
		mid = (l + r) >> 1
		if x >= mid*mid && x < (mid+1)*(mid+1) {
			return mid
		} else if x < mid*mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return 0
}

// 可以利用50题的思路，但实际上50题只有整数，没有小数操作
// x=res^2 求x
// 因为答案是整数，所以必然存在res^2<=x<(res+1)^2
// 暴力法如下，可以通过，但是速度很慢
//func mySqrt(x int) int {
//	if x == 0 { return 0 }
//	for i := 2; i <= x/2; i++ {
//		if x < (i+1)*(i+1) {
//			return i
//		}
//	}
//
//	return 1
//}
