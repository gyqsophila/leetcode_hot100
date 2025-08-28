package main

/*
 * 递归法，假设 list1 比 list 小，则 merge(list1,list2) 等价于 list1.Next=merge(list1.Next,list2)
 * 等价后的 merge 函数依然于之前有同样的结构
 * 最后返回一开始较小的 链表即可
 * 时间复杂度为 O(M+N)，空间复杂度为 O(M+N)
 */
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}
