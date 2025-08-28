package main

/*
 * 预制一个 preHead 节点作为合并后链表的头节点，用 prev 表示合并后链表的末尾节点
 * 迭代 2 个链表，谁的值更小就先将 prev 指向该链表的节点，并将该链表的指针后移
 * 最后会剩余一个一个链表的一个节点没有合并，谁不是 nil 就合并谁
 * 最后返回 preHead.Next 即可
 * 时间复杂度为 O(M+N)，空间复杂度为 O(1)
 */
func mergeTwoLists_1(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{Val: 0}
	// prev 始终指向合并链表的上一个节点
	prev := preHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return preHead.Next
}
