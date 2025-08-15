package main

/*
 * 主要思路是，计算每个字符串各个字符出现的次数，如果每个字母出现的次数都相同，那么变形后就是同一个字符串
 * 因为都是小写字母，用各个字母出现的次数作为 map 的key 来匹配，可以不需要排序，综合时间复杂度是 O(Nk)
 * 因为小写字母就 26 个，所以增加的时间复杂度可以忽略
 */
func groupAnagrams(strs []string) [][]string {
	timesMap := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, s := range str {
			key[s-'a']++
		}
		timesMap[key] = append(timesMap[key], str)
	}
	newStrs := make([][]string, 0, len(timesMap))
	for _, element := range timesMap {
		newStrs = append(newStrs, element)
	}
	return newStrs
}
