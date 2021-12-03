package main

func singleNumber(nums []int) int {
	bits := make([]uint32, 0)
	for i := 0; i < len(nums); i++ {
		n := uint32(nums[i])
		for j := 0; n != 0; j++ {
			bits[j] += n % 2
			n >>= 1
		}
	}

	res := uint32(0)
	for i := 0; i < len(bits); i++ {
		res += bits[i] % 3
		res <<= 1
	}

	return int(res)
}