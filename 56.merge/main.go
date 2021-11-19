package main

import (
	"fmt"
	"sort"
)

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

func main() {
	intervals := [][]int{
		{0,2}, {1,4}, {3,5},
	}
	fmt.Println(merge(intervals))
}

//输入：
//[[0,2],[1,4],[3,5]]
//输出：
//[[0,4],[3,5]]
//预期结果：
//[[0,5]]
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	// 对intervals进行排序，排序规则：按照tuple第0个元素的大小升序排列
	sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); {
		res = append(res, intervals[i])
		// 下一个跨到当前的边了
		if i + 1 < len(intervals) && intervals[i+1][0] <= res[len(res)-1][1] {
			// 把两个进行合并
			if intervals[i+1][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i+1][1]
			}
			i+=2
		} else {
			i++
		}
	}

	return res
}
