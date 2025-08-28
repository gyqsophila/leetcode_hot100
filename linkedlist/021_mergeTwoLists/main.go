package main

import (
	"fmt"
)

func main() {
	nums1 := []int{2}
	nums2 := []int{1}
	list1 := buildList(nums1)
	list2 := buildList(nums2)
	merged := mergeTwoLists(list1, list2)
	for merged != nil {
		fmt.Printf("%d ", merged.Val)
		merged = merged.Next
	}

}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 构造长度为 n 的链表
func buildList(nums []int) *ListNode {
	n := len(nums)
	var head, tail *ListNode
	for i := 0; i < n; i++ {
		node := &ListNode{Val: nums[i]}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}
