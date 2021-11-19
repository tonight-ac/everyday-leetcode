package main

//示例1:
//输入: a = "11", b = "1"
//输出: "100"
//示例2:
//输入: a = "1010", b = "1011"
//输出: "10101"

// 肯定还有更好的办法，这个是传统方案
func addBinary(a string, b string) string {
	res := ""

	for i, j, carry := len(a)-1, len(b)-1, 0; i >= 0 || j >= 0 || carry == 1; i, j = i-1, j-1 {
		v1, v2 := 0, 0
		if i >= 0 {
			v1 = int(a[i]-'0')
		}
		if j >= 0 {
			v2 = int(b[j]-'0')
		}

		tot := v1+v2+carry

		carry = tot/2

		res = string(rune(tot%2+'0')) + res
	}

	return res
}
