package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{1}
	// nums := []int{5, 4, -1, 7, 8}
	// nums := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	fmt.Println(maxSubArray(nums))
}
