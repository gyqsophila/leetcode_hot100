package main

/*
 * 暴力尝试，时间复杂度 O(N*N),性能很差
 */
func maxArea_force(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if cur := min(height[i], height[j]) * (j - i); cur > area {
				area = cur
			}
		}
	}
	return area
}
