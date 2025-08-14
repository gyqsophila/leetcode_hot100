package main

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

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	// result := twoSum([]int{3, 2, 4}, 6)
	// result := twoSum([]int{3, 3}, 6)
	println(result[0], result[1])
}
