package main

import (
	"fmt"
	"strings"
)

// zero
// one
// two
// three
// four
// five
// six
// seven
// eight
// nine
func main() {
	fmt.Println(originalDigits("owoztneoer"))
}
// TODO 11.24今日打卡
//var nums = [][]byte{
//	{'z','e','r','o'},     // z专属
//	{'o','n','e'},         //
//	{'t','w','o'},         // w专属
//	{'t','h','r','e','e'},
//	{'f','o','u','r'},
//	{'f','i','v','e'},
//	{'s','i','x'},         // x专属
//	{'s','e','v','e','n'},
//	{'e','i','g','h','t'},
//	{'n','i','n','e'},
//}

var nums = []map[byte]int{
	{'z':1,'e':1,'r':1,'o':1},
	{'o':1,'n':1,'e':1},
	{'t':1,'w':1,'o':1},
	{'t':1,'h':1,'r':1,'e':2},
	{'f':1,'o':1,'u':1,'r':1},
	{'f':1,'i':1,'v':1,'e':1},
	{'s':1,'i':1,'x':1},
	{'s':1,'e':2,'v':1,'n':1},
	{'e':1,'i':1,'g':1,'h':1,'t':1},
	{'n':2,'i':1,'e':1},
}

// 贪心存在多解，不行
func originalDigits(s string) string {
	// 放入bucket中，每个下标表示这个英文字母在s中出现的次数
	bucket := [26]int{}
	for _, v := range s { bucket[v-'a']++ }

	var b strings.Builder
	for i := 0; i < len(nums); i++ {
		times := 100000 // 1 <= s.length <= 105
		for k, v := range nums[i] {
			if v > bucket[k-'a'] {
				times = 0
				break
			}
			if times > bucket[k-'a']/v {
				times = bucket[k-'a']/v
			}
		}

		if times == 0 { continue }

		// 写入
		b.WriteString(strings.Repeat(string(rune(i+'0')), times))

		// 同时减去
		for k, v := range nums[i] {
			bucket[k-'a'] -= v * times
		}
	}

	return b.String()
}
