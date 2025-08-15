package main

/*
 * 用左右 2 个指针表示当前字符串的起始位置
 * 用 map 存储每个字母是否出现过
 * 当右移后的右指针所在字母重复，需要将左指针右移
 * 每次右指针移动后如果没有重复，说明不重复字符串长度+1
 * 比较后找到最大不重复字符串
 */
func lengthOfLongestSubstring_map(str string) int {
	n := len(str)
	if n <= 1 {
		return len(str)
	}
	charMap := map[byte]int{}
	left := 0
	right := 0
	maxLength := 0
	for right < n {
		if charMap[str[right]] > 0 {
			delete(charMap, str[left])
			left++
		} else {
			charMap[str[right]] = 1
			maxLength = max(maxLength, right-left+1)
			right++
		}
	}
	return maxLength
}
