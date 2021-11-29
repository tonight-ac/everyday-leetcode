package main

// 肯定还有更好的办法，这个是传统方案
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits)-1; carry == 1 && i >= 0; i-- {
		digits[i], carry = (digits[i]+carry)%10, (digits[i]+carry)/10
	}

	// 如果最后多进出来一位，需要插入在数列首位
	if carry == 1 {
		// 记录一下末尾数字，因为等下整体后移会把最后一位盖掉
		end := digits[len(digits)-1]
		for i := 1; i < len(digits); i++ {
			digits[i] = digits[i-1]
		}
		// 添加进位
		digits[0] = carry
		// 把原来的末尾添加回来
		digits = append(digits, end)
	}

	return digits
}
