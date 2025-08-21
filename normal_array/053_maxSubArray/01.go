package main

/*
 * 贪心算法
 * 核心思想是，如果包含当前位置之前的 sum 小于当前位置的值，则直接丢弃当前位置之前的数
 * 否则累加到最大值
 */
func maxSubArray_01(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], curSum+nums[i])
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
