package main

import (
	"fmt"
	"slices"
)

func main() {
	// internals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	internals := [][]int{{1, 4}, {4, 5}}
	for _, cell := range merge(internals) {
		fmt.Printf("%v ", cell)
	}
}

/*
 * 先按照数组左端点排序，然后能够合并的区间一定是连续的
 * 然后区间右端点比较和替换即可
 * 主要时间开销在排序上， O(nlogn)
 */
func merge(intervals [][]int) [][]int {
	// 根据左端点排序
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	result := [][]int{intervals[0]}
	resultIndex := 0
	for i := 1; i < len(intervals); i++ {
		// 新的区间
		if intervals[i][0] > result[resultIndex][1] {
			result = append(result, intervals[i])
			resultIndex++
			continue
		}
		// 右边界拓展
		if intervals[i][1] > result[resultIndex][1] {
			result[resultIndex][1] = intervals[i][1]
		}
	}
	return result
}
