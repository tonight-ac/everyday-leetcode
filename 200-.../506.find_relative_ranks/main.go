package main

import (
	"sort"
	"strconv"
)

var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	n := len(score)
	type pair struct{ score, idx int }
	arr := make([]pair, n)
	for i, s := range score {
		arr[i] = pair{s, i}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].score > arr[j].score })

	ans := make([]string, n)
	for i, p := range arr {
		if i < 3 {
			ans[p.idx] = desc[i]
		} else {
			ans[p.idx] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
