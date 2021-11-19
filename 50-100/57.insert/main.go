package main

//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10]重叠。
// TODO 未完成
func insert(intervals [][]int, newInterval []int) [][]int {
	// 两步走，先把newInterval根据newInterval[0]大小给塞进去，然后往两边扩展
	// 寻找插入点
	idx := len(intervals)
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[0] {
			idx = i
			break
		}
	}

	left, right := idx, idx
	// 左边向左寻找可能merge的interval
	for left >= 0 {

	}

	// 右边向右寻找可能merge的interval
	for right < len(intervals) {

	}

	return nil
}
