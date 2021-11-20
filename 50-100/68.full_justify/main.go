package main

import "strings"

//输入：
//["This", "is", "an", "example", "of", "text", "justification."]
//16
//输出：
//["This    is    an","of text         "]
//预期结果：
//["This    is    an","example  of text","justification.  "]
// 先确定分组，再确定' '的个数
func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	row := make([]string, 0)
	width := maxWidth
	wordsWidth := 0

	for i := 0; i < len(words); i++ {
		// 还可以贪心放入
		if len(words[i]) <= width {
			width -= len(words[i])+1
			wordsWidth += len(words[i])
			row = append(row, words[i])
		} else { // 不能放入更多，直接置width为0
			width = 0
		}

		// 开始处理这一行
		if width <= 0 {
			if i < len(words) - 1 {
				// 不是最后一行，左右两端对齐
				spaceWidth := maxWidth-wordsWidth
				div, mod := spaceWidth/(len(row)-1), spaceWidth%(len(row)-1)
				s := ""
				for i := 0; i < len(row); i++ {
					if i != len(row) - 1 {
						s += row[i] + strings.Repeat(" ", div)
						if i < mod {
							s += " "
						}
					} else {
						s += row[i]
					}
				}
				res = append(res, s)
			} else {
				// words已经遍历结束目前为最后一行需要左对齐
				s := strings.Join(row, " ")
				// 补齐空格
				s = s + strings.Repeat(" ", maxWidth-len(s))
				res = append(res, s)
			}

			// 开始下一行，条件清空
			row = make([]string, 0)
			width = maxWidth
			wordsWidth = 0
		}
	}

	return res
}
