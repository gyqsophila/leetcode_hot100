package main

/*
 * 动态规划
 * 核心思想是，如果前一只数字的值大于 0，则累加到当前位置
 * 否则累加到最大值
 */
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxSum = max(nums[i], maxSum)
	}
	return maxSum
}
