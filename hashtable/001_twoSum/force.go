package main

/*
 * 暴力匹配，没啥好说的，只需要往后匹配就行了
 */

func twoSum_force(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
