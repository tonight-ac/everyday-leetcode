package main

import (
	"fmt"
	"strings"
)

// 383的复杂不去重版本
// leetcode [leet, code]


//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
func main()  {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
}

// TODO 超时了
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 { return true }
	for _, v := range wordDict {
		if idx := strings.Index(s, v); idx != -1 {
			ok1 := wordBreak(s[:idx], wordDict)
			ok2 := wordBreak(s[idx+len(v):], wordDict)
			if ok1 && ok2 {
				return true
			}
		}
	}
	return false
}

