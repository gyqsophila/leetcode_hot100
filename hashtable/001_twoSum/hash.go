package main

/*
 * 遍历过的数字存在 map 里，用新的数字在 map 找能组合成目标值的数字，只需要遍历一遍
 */
func twoSum(nums []int, target int) []int {
	// num - index
	existedNums := make(map[int]int)
	for i, x := range nums {
		if j, ok := existedNums[target-x]; ok {
			return []int{i, j}
		}
		existedNums[x] = i
	}
	return nil
}
