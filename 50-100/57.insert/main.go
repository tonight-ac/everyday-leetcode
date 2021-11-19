package main

import "fmt"

//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10]重叠。

func main() {
	// 使用下图的方式定位newInterval两个点的位置，奇数为空档，偶数为点，偶数除2还是偶数为[0]，奇数为[1]
	// 1    2    3    5
	// 0 1* 2 3* 4 5* 6
	fmt.Println(insert([][]int{{1,5}}, []int{2,3}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 { return [][]int{newInterval} }
	// 两步走，先把newInterval根据newInterval[0]大小给塞进去，然后往两边扩展

	res := make([][]int, 0)

	// 寻找插入点，先插入
	for i := 0; i < len(intervals); i++ {
		// 定位左边界限
		if intervals[i][0] > newInterval[0] {
			// 因为intervals[i][0] > newInterval[0]，所以newInterval[0] >= intervals[i-1][0]
			// 所以newInterval[0]和intervals[i-1]可能存在合并关系
			// 所以把intervals[i-1]放入
			if i > 0 {
				res = append(res, intervals[0:i]...)
				// 多切一个给newInterval用，不用额外申请空间了
				intervals = intervals[i-1:]
				intervals[0] = newInterval
			} else {
				// 复用56的merge算法
				// res[len(res)-1]为待处理项
				res = append(res, newInterval)
			}
			break
		}
	}

	// newInterval[0]大于所有intervals[i][0]
	// 把intervals放入res中，newInterval放入intervals作为待处理tuple
	if len(res) == 0 {
		res = intervals
		intervals = [][]int{newInterval}
	}

	// 再从插入点之后开始merge
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			// 没跨到边，直接下一个，不需要合并
			res = append(res, intervals[i])
		} else if intervals[i][1] > res[len(res)-1][1] {
			// 跨上边了，并且是不含与的情况，需要更新右边界
			res[len(res)-1][1] = intervals[i][1]
		}
	}

	return res
}
