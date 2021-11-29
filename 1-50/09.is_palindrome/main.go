package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xx, yy := x, 0

	for ; xx != 0; xx/=10 {
		yy = yy * 10 + xx % 10
	}

	return x == yy
}