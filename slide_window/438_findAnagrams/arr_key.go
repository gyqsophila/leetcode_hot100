package main

/*
 * 因为都是小写字母，可以用 [26]int{} 数组来存储各个字母出现的次数
 * 只要出现的次数的字母想同，则必然是异维词
 * 维护更新 s 当前字符串的 key 数组，可以无需重复申请内存
 */
func findAnagrams(s string, p string) []int {
	result := []int{}
	lenS := len(s)
	lenP := len(p)
	pkey := [26]int{}
	for i := 0; i < lenP; i++ {
		pkey[p[i]-'a']++
	}
	sKey := [26]int{}
	for i := 0; i < lenS; i++ {
		sKey[s[i]-'a']++
		if i >= lenP-1 {
			if sKey == pkey {
				result = append(result, i-(lenP-1))
			}
			// 删除开始的字母的计数，往右移
			sKey[s[i-(lenP-1)]-'a']--
		}
	}
	return result
}
