package main

import "fmt"

func main() {
	//fmt.Println(INT32MAX)
	//fmt.Println(INT32MIN)
	fmt.Println(myAtoi("42"))
}

//读入字符串并丢弃无用的前导空格
//检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
//读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
//如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
//返回整数作为最终结果。

const INT32MAX = 2147483647
const INT32MIN = -2147483648

func myAtoi(s string) int {
	idx := 0
	char := int32(1)
	res := int32(0)

	// 去除空格
	for ; idx < len(s); idx++ {
		if s[idx] != ' ' {
			break
		}
	}

	// 避免越界
	if idx >= len(s) {
		return 0
	}

	// 去除符号
	if s[idx] == '-' {
		char = -1
		idx++
	} else if s[idx] == '+' {
		char = 1
		idx++
	}

	// 整理数字
	for ; idx < len(s); idx++ {
		if s[idx] >= '0' && s[idx] <= '9' {
			temp := res*10+char*int32(s[idx]-'0')
			// 进行可逆操作判断是否发生越界
			if res != temp/10 {
				if char == 1 {
					return INT32MAX
				}
				return INT32MIN
			}
			res = temp
		} else { // 非数字立刻截断，返回在此之前读出的数字结果
			break
		}
	}

	return int(res)
}
