package main

import (
	"fmt"
)

func main() {
	// nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	nums, k := []int{1, 2, 3, 4, 5, 6}, 1
	// nums, k := []int{-1, -100, 3, 99}, 2
	fmt.Println(rotate(nums, k))
}

/*
 * 用新的数组存储轮转后的结果，再拷贝回原数组，只是用到了 go 自带的 append 方法
 */
func rotate_append(nums []int, k int) []int {
	length := len(nums)
	k %= length
	newNums := append(nums[length-k:], nums[0:length-k]...)
	copy(nums, newNums)
	return nums
}

/*
 * 用新的数组存储轮转后的结果，再拷贝回原数组
 */
func rotate_new(nums []int, k int) []int {
	length := len(nums)
	newArray := make([]int, length)
	for i := 0; i < length; i++ {
		newPos := (i + k) % length
		newArray[newPos] = nums[i]
	}
	copy(nums, newArray)
	return nums
}

/*
 * 以一个临时变量存储要移动的元素，找到要移动的目标位置，交换临时元素和目标元素，在移动 k 个位置
 * 重复者个过程直到位置回到开始的元素位置
 * 要遍历的轮数就是 n 和 k 的最小公倍数，每轮的开始位置都要在上一轮开始的位置上+1
 */
func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n
	// start 表示当前一轮遍历的开始元素, loopCount 表示要遍历的轮数,美轮遍历以回到 start 元素为结束
	for start, loopCount := 0, gcd(k, n); start < loopCount; start++ {
		// tmp 存储当前被覆盖的元素，cur 表示当前轮在遍历的元素索引
		tmp, cur := nums[start], start
		// 刚开始 cur 必然等于 start，遍历开始后再次遇到相等则退出遍历
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], tmp, cur = tmp, nums[next], next
		}
	}
	return nums
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
