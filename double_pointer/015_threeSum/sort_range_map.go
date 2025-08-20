package main

import (
	"sort"
)

/*
 * 对数组排序后，问题会简化，必然存在 x<y<z
 * 在遍历 x,y 后在 map 中寻找 z，还需要满足 x<y<z
 */
func threeSum_map(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)

	numsMap := make(map[int][]int)
	// 放进 map 同时也能去重
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = append(numsMap[nums[i]], i)
	}
	for x := 0; x < len(nums)-2; x++ {
		// 跳过重复的 x
		if x > 0 && nums[x-1] == nums[x] {
			continue
		}
		for y := x + 1; y < len(nums)-1; y++ {
			// 跳过重复的 y
			if y > x+1 && nums[y] == nums[y-1] {
				continue
			}
			// 寻找最后一个值 z
			if zList, ok := numsMap[0-nums[x]-nums[y]]; ok {
				for _, z := range zList {
					if z > x && z > y {
						results = append(results, []int{nums[x], nums[y], nums[z]})
						// 当 x,y 固定后，z 也只能出现一次，否则会重复
						break
					}
				}
			}
		}
	}
	return results
}
