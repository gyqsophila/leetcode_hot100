package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	R := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= R
		R *= nums[i]
	}
	return result
}

/*
 * 先分别计算索引位置左边元素之乘积与右边元素之乘积，分别存储在 2 个数组中，这样至需要遍历 2 遍
 * 在遍历一遍，用索引左边之乘积 * 索引右边之乘积，就是除自身之外的元素之乘积
 * 由于排 0 的左边和 n-1 的右边都没有元素，所以值应该是 1，这样不会改变另外一边乘积的值
 * 时间负责度 0(N),空间复杂度O(N)
 */
func productExceptSelf_on(nums []int) []int {
	n := len(nums)
	L, R, result := make([]int, n), make([]int, n), make([]int, n)
	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	R[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}
	for i := 0; i < n; i++ {
		result[i] = L[i] * R[i]
	}
	return result
}

/*
 * 双循环暴力计算，会超时
 */
func productExceptSelf_force(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				sum *= nums[j]
			}
		}
		result[i] = sum
	}
	return result
}
