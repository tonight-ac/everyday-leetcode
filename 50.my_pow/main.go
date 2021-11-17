package main

import "fmt"

func main() {
	fmt.Println(myPow(2.1000, 3))
}

func myPow(x float64, n int) float64 {
	if n == 0 { return 1.0 }
	res := x
	if n > 0 {
		for i := 0; i < n-1; i++ {
			res *= x
		}
	} else {
		for i := 0; i < -(n-1); i++ {
			res /= x
		}
	}

	return res
}
