package main

var start int
var res int
func maxPower(s string) int {
	start = 0
	res = 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			start = i
		} else {
			if res < i - start + 1 {
				res = i - start + 1
			}
		}
	}

	return res
}