package main

import "math/rand"

type Solution struct {
	m int
	n int
	size int
	uni map[int]bool
}

func Constructor(m int, n int) Solution {
	return Solution{
		m: m,
		n: n,
		size: m*n,
		uni: map[int]bool{},
	}
}

func (s *Solution) Flip() []int {
	for {
		tar := rand.Intn(s.size)
		if !s.uni[tar] {
			s.uni[tar] = true
			return []int{tar/s.n, tar%s.n}
		}
	}
}

func (s *Solution) Reset()  {
	s.uni = map[int]bool{}
}
