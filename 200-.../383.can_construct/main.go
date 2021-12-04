package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) { return false }
	m := make(map[uint8]int16)
	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]]++
	}

	count := len(ransomNote)
	for i := 0; i < len(magazine) && count != 0; i++ {
		if v, ok := m[magazine[i]]; ok && v > 0 {
			count--
			m[magazine[i]]--
		}
	}

	return count == 0
}
// 在magazine中搜索ransomNote
//func canConstruct(ransomNote string, magazine string) bool {
//	m := make(map[uint8]int16)
//	for i := 0; i < len(magazine); i++ {
//		m[magazine[i]]++
//	}
//
//	for i := 0; i < len(ransomNote); i++ {
//		if m[ransomNote[i]] < 1 {
//			return false
//		}
//		m[ransomNote[i]]--
//	}
//
//	return true
//}