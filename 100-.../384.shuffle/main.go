package main

import (
	"math/rand"
)

type Solution struct {
	n []int
}

func Constructor(nums []int) Solution {
	return Solution{
		n: nums,
	}
}

func (s *Solution) Reset() []int {
	return s.n
}

func (s *Solution) Shuffle() []int {
	res := make([]int, len(s.n))
	copy(res, s.n)

	// 不要这句，可以提升速度
	//rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(res); i++ {
		tar := rand.Intn(len(res))
		res[i], res[tar] = res[tar], res[i]
	}

	return res
}
