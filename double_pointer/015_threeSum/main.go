package main

import "fmt"

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 1, 1}
	nums := []int{0, 0, 0}
	// nums := []int{0, 0, 0, 0}
	for _, tuple := range threeSum(nums) {
		fmt.Printf("%v\n", tuple)
	}
}
