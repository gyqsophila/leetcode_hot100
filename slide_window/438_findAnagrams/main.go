package main

import "fmt"

func main() {
	sArray := []string{
		"cbaebabacd",
		"abab",
		"aa",
	}
	pArray := []string{
		"abc",
		"ab",
		"bb",
	}

	for i := range sArray {
		result := findAnagrams(sArray[i], pArray[i])
		fmt.Println(result)
	}
}
