package main

func convert(s string, numRows int) string {
	// 1       9 2 * numRows - 2
	// 2     8 1
	// 3   7
	// 4 6
	// 5
	cycle := 2 * numRows - 2
	res := ""

	for row := 0; row < numRows; row++ {
		diff := cycle - 2 * numRows
		for idx := row; idx < len(s); {
			res += string(s[idx])
			idx += diff
			if diff != cycle {
				diff = cycle - diff
			}
		}
	}

	return res
}
