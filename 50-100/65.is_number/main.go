package main

import (
	"fmt"
)

func main()  {
	fmt.Println(isNumber("."))
}
//输入：
//".-4"
//输出：
//true
//预期结果：
//false
func isNumber(s string) bool {
	// 先定位首个'e'/'E'
	// 然后分别对'e'/'E'两侧的数字做判断
	for i := 0; i < len(s); i++ {
		if s[i] == 'e' || s[i] == 'E' {
			return isFloat(s[:i]) && isInt(s[i+1:])
		}
	}

	return isFloat(s)
}

// isDigit point表示是否允许小数
func isFloat(s string) bool {
	if len(s) == 0 { return false }

	idx := 0
	// 先判断符号
	if s[idx] == '-' || s[idx] == '+' { idx++ }
	// 如果只有符号，不是数字
	if idx == len(s) { return false }

	for ; idx < len(s); idx++ {
		if s[idx] == '.' {
			if len(s[:idx]) == 0 {
				return isInt(s[idx+1:])
			} else if len(s[idx+1:]) == 0 {
				return isInt(s[:idx])
			}
			return isInt(s[:idx]) && isInt(s[idx+1:])
		}
	}

	return isInt(s)
}

// isDigit point表示是否允许小数
func isInt(s string) bool {
	if len(s) == 0 { return false }

	idx := 0
	// 先判断符号
	if s[idx] == '-' || s[idx] == '+' { idx++ }
	// 如果只有符号，不是数字
	if idx == len(s) { return false }

	for ; idx < len(s); idx++ {
		if s[idx] < '0' || s[idx] > '9' {
			return false
		}
	}

	return true
}


// TODO 未完成
// 类似第8题 不需要特别精简，可读性第一
// 下面这种方式未通过，且特殊case太多，漏洞堵不上，换一种方式
//func isNumber(s string) bool {
//	idx := 0
//	// 先判断符号
//	if s[idx] == '-' || s[idx] == '+' { idx++ }
//	// 如果只有符号，不是数字
//	if idx == len(s) { return false }
//
//	// 开始判断数字
//	for ; idx < len(s); idx++ {
//		if s[idx] == '.' {
//			break
//		} else if s[idx] == 'e' || s[idx] == 'E' {
//			break
//		} else if s[idx] < '0' || s[idx] > '9' {
//			return false
//		}
//	}
//
//	// 遍历结束了，纯数字
//	if idx == len(s) { return true }
//
//	// 判断是否有小数
//	if s[idx] == '.' {
//		// '.'前面后面都可以没东西，表示省略0，但前后不能都没有数字
//		//if idx == 0 || idx == len(s) - 1 { return false }
//		if (idx-1<0 || s[idx-1] < '0' || s[idx-1] > '9') && (idx+1>=len(s) || (s[idx+1] < '0' || s[idx+1] > '9')) {
//			return false
//		}
//		idx++ // 跳过'.'
//		for ; idx < len(s); idx++ {
//			if s[idx] == 'e' || s[idx] == 'E' {
//				break
//			} else if s[idx] < '0' || s[idx] > '9' {
//				return false
//			}
//		}
//	}
//
//	// 遍历结束了，含小数纯数字
//	if idx == len(s) { return true }
//
//	// 科学计数法
//	if s[idx] == 'e' || s[idx] == 'E' {
//		// 'e'/'E'前面后面都需要有东西
//		if idx == 0 || idx == len(s) - 1 { return false }
//		idx++ // 跳过'e'/'E'
//		for ; idx < len(s); idx++ {
//			if s[idx] < '0' || s[idx] > '9' {
//				return false
//			}
//		}
//	}
//
//	return true
//}
