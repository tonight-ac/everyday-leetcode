package main

//输入：s = "mississippi" p = "mis*is*p*."
//输出：false

// 讨巧的递归做法
func isMatch(s string, p string) bool {
	// 如果p已经匹配完了，那就判断一下s有没有匹配完
	if p == "" {
		return s == ""
	}

	// 判断第一个字符是否匹配
	flag := s != "" && (s[0] == p[0] || p[0] == '.')

	// 判断p未来是否存在* 提前处理掉
	if len(p) > 1 && p[1] == '*' {
		// 跳过按空处理 或者 匹配掉向后继续匹配
		return isMatch(s, p[2:]) || (flag && isMatch(s[1:], p))
	}

	// 不存在通配符，直接向后处理
	return flag && isMatch(s[1:], p[1:])
}
