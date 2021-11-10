package main

//输入：s = "mississippi" p = "mis*is*p*."
//输出：false

func isMatch(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}

	for i, j := 0, 0; i < len(s) || j < len(p); {
		if s[i] == p[j] {

		}
	}

	return true
}