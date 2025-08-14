package main

/**
 * 将数组存储在 map 中，依次寻找不存在前序数字的连续数字串的长度
 * 每个连续数字串结束后，记录数字串的长度并与历史最长串比较
 * 这样就能找到最长的连续数字串的长度
 * 综合时间复杂度是 O(N) 因为若没有下一个连续数字，内循环就会停止，内循环搜索次数不会超过元素个数的 2倍，因此 2N 也是 N
 */

func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}
	longestLen := 0
	for num := range numsMap {
		// 如果不存在前序数字，才需要计算当前连续数字的长度
		if !numsMap[num-1] {
			currentLongLen := 1
			currentNum := num
			// 若存在后序数字则继续搜索
			for numsMap[currentNum+1] {
				currentNum++
				currentLongLen++
			}
			// 当前连续数字结束后与历史最长比较
			if currentLongLen > longestLen {
				longestLen = currentLongLen
			}
		}
	}
	return longestLen
}
