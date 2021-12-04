package main

import "fmt"

func main() {
	n := -1
	fmt.Println(n<<1)
	//fmt.Println(singleNumber([]int{-1,-1,-2,-1}))
}

func singleNumber(nums []int) int {
	bits := make([]int32, 32)

	for i := 0; i < 32; i++ {
		bits[i] = ^(1 << i)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(bits); j++ {
			n := int32(nums[i])
			n &= bits[j]
			bits[j] ^= n
		}
	}

	//res := 0
	//for i := 0; i < len(bits); i++ {
	//	res += (bits[i] % 3) << i
	//}

	//if bits[len(bits)-1] % 3 == 0 {
	//	return int(res)
	//}
	//return -int(res)
	return 0
}