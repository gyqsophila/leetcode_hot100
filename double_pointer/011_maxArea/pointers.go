package main

/*
 * 用 left, right 表示容器的左右边界所在位置
 * 不断比较左右指针所组成容器的容积
 * 因为容器的容积是有最小的边界和左右的距离决定的
 * 因此每次仅移动较小的那个边界，这样可能在距离 -1 的同时增加最小的边界的值
 * 而如果移动较大的边界，则会导致距离减小的同时，也必然会降低容器的可能的最大容积，不是我们的目的
 */

func maxArea(height []int) int {
	left, right, area := 0, len(height)-1, 0
	for left < right {
		if cur := min(height[left], height[right]) * (right - left); cur > area {
			area = cur
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
