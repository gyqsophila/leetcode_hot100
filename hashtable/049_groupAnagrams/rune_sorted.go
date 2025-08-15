package main

import (
	"sort"
	"strings"
)

/*
 * 主要思路是，将每个字符串按照字母排序，排序后的字母如果相同则可以分为同一组，并存储在 map 中
 * 每个字符穿只需要扫描一便，但是字符串排序会增加时间复杂度,综合时间复杂度为 O(Nklogk)
 */

func groupAnagrams_bak(strs []string) [][]string {
	sortedExistMap := map[string][]string{}
	for _, str := range strs {
		key := strResort(str)
		sortedExistMap[key] = append(sortedExistMap[key], str)
	}
	newStrings := [][]string{}
	for _, strs := range sortedExistMap {
		newStrings = append(newStrings, strs)
	}
	return newStrings
}

func strResort(str string) string {
	newStr := []string{}
	for _, s := range str {
		newStr = append(newStr, string(s))
	}
	sort.Strings(newStr)
	return strings.Join(newStr, "")
}
