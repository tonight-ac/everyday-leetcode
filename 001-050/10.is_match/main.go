package main

//输入：s = "mississippi" p = "mis*is*p*."
//输出：false

// 直接相同
// 只有 .
// 有 * ，前面是普通字符
// 有 * ，前面是 .

// s: mississippi
// p: mis*is*p*.
func isMatch(s string, p string) bool {
	// s为空 p为空 为真
	// s不为空 p不为空 不能判断
	// s为空 p不为空 不能判断真假 ex .*
	// s不为空 p为空 为假
	if len(p) == 0 {
		return len(s) == 0
	}

	// 拆解为子问题，判断第一个字符
	firstMatch := s != "" && (s[0] == p[0] || p[0] == '.')

	// 判断p未来是否存在* 提前处理掉
	if len(p) > 1 && p[1] == '*' {
		// 跳过按空处理 或者 匹配掉向后继续匹配
		return (firstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}

	// 不存在通配符，直接向后处理
	return firstMatch && isMatch(s[1:], p[1:])
}
