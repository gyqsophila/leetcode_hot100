package main

/*
 * 思路是一样的，但是用数组来存储字母转数字后的值
 * 按照题意，出现的字母的 ascll 码的值都在 0～128 以内
 * 用一个 128 长度的 int 数组足够存储，判断重复会比 map 更快（map 的 key 会有碰撞计算)
 */
func lengthOfLongestSubstring(str string) int {
	n := len(str)
	if n <= 1 {
		return n
	}
	asc := [128]int{}
	left, right := 0, 0
	maxLength := 0
	for right < n {
		asc[str[right]]++
		for asc[str[right]] > 1 {
			asc[str[left]]--
			left++
		}
		maxLength = max(maxLength, right-left+1)
		right++
	}
	return maxLength
}
