package main

/*
 * left 为已处理的非 0 数字的最后一位的索引
 * right 为待处理的的数字最左边一位的索引
 * 那么，left 左边都是处理好的非 0 数字，left 和 right 之间都是 0
 * 每次左右指针交换值，都是把 right 处的非 0 数字与 left 的 0 交换
 * 因为不存在非 0 数字之间的交换，所以非 0 数字依然保持之间的顺序
 */
func moveZeroes(nums []int) {
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
