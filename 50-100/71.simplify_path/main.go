package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
}

// 用一个双向队列来处理问题
func simplifyPath(path string) string {
	stack := make([]string, 0)
	for i := 0; i < len(path); {
		j := i+1
		for ; j < len(path); j++ {
			if path[j] == '/' { break }
		}
		p := path[i:j]

		if p == "/.." {
			if len(stack) > 0 {
				// 返回上一级目录，即删除最后一个
				stack = stack[:len(stack)-1]
			}
		} else if p != "/" && p != "/." {
			stack = append(stack, p)
		}

		// 更新i到下一个'/'处
		i = j
	}

	if len(stack) == 0 {
		return "/"
	}

	return strings.Join(stack, "")
}
