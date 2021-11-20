package main

import (
	"fmt"
	"strings"
)

//输入：
//["This", "is", "an", "example", "of", "text", "justification."]
//16
//输出：
//["This    is    an","of text         "]
//预期结果：
//["This    is    an","example  of text","justification.  "]

func main() {
	res := fullJustify([]string{"What","must","be","acknowledgment","shall","be"}, 16)
	fmt.Println(res)
	for _, v := range res {
		fmt.Println(v)
	}
}

// 先确定分组，再确定' '的个数
func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	row := make([]string, 0)
	width := maxWidth
	wordsWidth := 0

	for i := 0; i < len(words); {
		// 还可以贪心放入
		if len(words[i]) <= width {
			width -= len(words[i])+1
			wordsWidth += len(words[i])
			row = append(row, words[i])
			i++
		} else { // 不能放入更多，直接置width为0
			width = 0
		}

		// 开始处理这一行
		if width <= 0 {
			//s := ""
			var b strings.Builder
			if len(row) > 1 {
				// 不是最后一行，左右两端对齐
				spaceWidth := maxWidth-wordsWidth
				div, mod := spaceWidth/(len(row)-1), spaceWidth%(len(row)-1)
				for j := 0; j < len(row); j++ {
					if j != len(row) - 1 {
						b.WriteString(row[j])
						// 有求尽量平均
						b.WriteString(strings.Repeat(" ", div))
						//s += row[j] + strings.Repeat(" ", div)
						if j < mod { // 存在余数，需要多加一个" "，因为要求左边空格多余右侧
							b.WriteString(" ")
							//s += " "
						}
					} else {
						b.WriteString(row[j])
						//s += row[j]
					}
				}
			} else {
				b.WriteString(row[0])
				b.WriteString(strings.Repeat(" ", maxWidth-len(row[0])))
				//s = row[0] + strings.Repeat(" ", maxWidth-len(row[0]))
			}

			res = append(res, b.String())

			// 开始下一行，条件清空
			row = make([]string, 0)
			width = maxWidth
			wordsWidth = 0
		}
	}

	// 处理最后一行，左对齐
	// words已经遍历结束目前为最后一行需要左对齐
	if len(row) > 0 {
		s := strings.Join(row, " ")
		// 补齐空格
		s = s + strings.Repeat(" ", maxWidth-len(s))
		res = append(res, s)
	}

	return res
}
