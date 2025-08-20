package main

/*
 * 用一个指针表示非 0 数字应该排列的位置，遍历数据遇到非 0 数字则依次放置在对应的位置上
 * 同时统计非 0 数字的个数，在遍历一遍和放置非 0 数字后，将数组后续位置都修改为 0 即可
 */
func moveZeroes_single_pointer(nums []int) {
	numIndex := 0
	numsCounter := 0
	for i := range nums {
		if nums[i] != 0 {
			numsCounter++
			nums[numIndex] = nums[i]
			numIndex++
		}
	}
	for i := numsCounter; i < len(nums); i++ {
		nums[i] = 0
	}
}
