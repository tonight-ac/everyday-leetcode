package main

import "fmt"

func titleToNumber(columnTitle string) int {
	if len(columnTitle) == 0 {
		return 0
	}

	res := 0
	factor := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += int(columnTitle[i] - 'A' + 1) * factor
		factor *= 26
	}

	return res
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}