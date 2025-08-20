package main

import (
	"sort"
)

/*
 * 对数组排序后，问题会简化，必然存在 x<y<z
 * 在遍历 x 后，用 y 和 z 表示 2 个数的索引，则问题为 x+y+z=0 && x<y<z
 */
func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)

	length := len(nums)
	for x := 0; x < length; x++ {
		// 跳过重复的 x
		if x > 0 && nums[x-1] == nums[x] {
			continue
		}
		z := length - 1
		for y := x + 1; y < length; y++ {
			// 跳过重复的 y
			if y > x+1 && nums[y] == nums[y-1] {
				continue
			}
			for y < z && nums[y]+nums[z] > -(nums[x]) {
				z--
			}

			if y == z {
				break
			}

			if nums[x]+nums[y]+nums[z] == 0 {
				results = append(results, []int{nums[x], nums[y], nums[z]})
			}
		}
	}
	return results
}
