package main

import "fmt"

func main() {
	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "*aa*ba*a*bb*aa*ab*a*aaaaaa*a*aaaa*bbabb*b*b*aaaaaaaaa*a*ba*bbb*a*ba*bb*bb*a*b*bb"))
}
//"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb"
//"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"
// 还需要多看，多理解几遍
// 此题解系官方的贪心题解，为何这种情况下不会超时呢？
// 核心思路：将p等价为：*v1*v2*.....*vn*和s匹配的问题
// s中间只要出现v1......vn即可说明匹配成功
// 但是p的末尾可能不是*的情况，这就说明需要和p的末尾进行严格匹配，所以需要先将问题化简为以上问题，需要先从末尾进行处理和匹配
func isMatch(s string, p string) bool {
	// 从后向前匹配，去尾，只匹配完全相同的，相等或p[:len(p)-1]=='?'
	// 遇到'*'停止，出现不相同的直接返回false
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if charMatch(s[len(s)-1], p[len(p)-1]) {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	// 如果p完全匹配完了，说明不包含'*'
	// 如果s也没有了，则说明s和p完全匹配
	if len(p) == 0 {
		return len(s) == 0
	}

	sIndex, pIndex := 0, 0
	sRecord, pRecord := -1, -1
	// 开始从前向后扫，注意条件，是pRecord < len(p)而非pIndex < len(p)
	for sIndex < len(s) && pRecord < len(p) {
		if p[pIndex] == '*' { // 如果p是一个'*'
			pIndex++ // 直接跳过p，意思是当这个'*'不存在，直接跳过
			sRecord, pRecord = sIndex, pIndex // 并且记录这一次出现'*'的位置
		} else if charMatch(s[sIndex], p[pIndex]) { // 如果不是'*' 则看是不是完全相同
			sIndex++ // s指向下一个
			pIndex++ // p指向下一个
		} else if sRecord != -1 && sRecord + 1 < len(s) { // 存在上一个'*'的记录，可以回退，并且s不是最后一个，因为如果是最后一个，那必然不可能到这个循环处，因为p多了一个，那就应该返回false
			sRecord++ // 指向s下一个，意思是这个被'*'给匹配吃掉了
			sIndex, pIndex = sRecord, pRecord // p回退到上一个'*'处重新遍历，回退到第一次遇到'*'的下一个
		} else { // 不符合以上条件
			return false
		}
	}

	// 检查剩余的p是不是纯*，或者空串
	return allStars(p, pIndex, len(p))
}

func allStars(str string, left, right int) bool {
	for i := left; i < right; i++ {
		if str[i] != '*' {
			return false
		}
	}
	return true
}

func charMatch(u, v byte) bool {
	return u == v || v == '?'
}

//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。

//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。

// 类似第10题的解法，不过递归超时了
// s的长度会大幅制约算法性能
// 死活超时，需要别的办法
//func isMatch(s string, p string) bool {
//	// 如果p已经匹配完了，那就判断一下s有没有匹配完
//	if p == "" {
//		return s == ""
//	}
//
//	if p[0] == '*' {
//		// 清洗一下多余的*
//		for i := 0; i < len(p) - 1; i++ {
//			if p[i+1] != '*' {
//				p = p[i:]
//				break
//			}
//		}
//
//		// 如果s[0]和p[1]不相同，则为了匹配，必然需要让*吃掉s[0]
//		if len(s) > 0 && len(p) > 1 {
//			for i := 0; i < len(s); i++ {
//				if s[i] == p[1] {
//					s = s[i:]
//					break
//				}
//			}
//		}
//
//		// p为'*' 那分为匹配当前s和不匹配当前s两种结果
//		return (len(s) > 0 && isMatch(s[1:], p)) || isMatch(s, p[1:])
//	}
//
//	// 相等 或者 p为'?'可以匹配任何单个字符
//	return len(s) > 0 && (s[0] == p[0] || p[0] == '?') && isMatch(s[1:], p[1:])
//}
