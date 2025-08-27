package main

import "fmt"

func main() {
	n := 100000000 // 改这里测试不同长度，比如 10000 / 50000 / 100000
	head := buildList(n)

	fmt.Printf("List built with %d nodes\n", n)

	// 递归翻转（可能栈溢出）
	fmt.Println("Start recursive reverse...")
	newHead := reverseList(head) // 大链表会 panic: stack overflow
	fmt.Println("Recursive reverse finished. Head:", newHead.Val)

	// 迭代翻转（安全）
	// fmt.Println("Start iterative reverse...")
	// newHead2 := reverseListIter(head)
	// fmt.Println("Iterative reverse finished. Head:", newHead2.Val)
}

// 构造长度为 n 的链表
func buildList(n int) *ListNode {
	var head, tail *ListNode
	for i := 0; i < n; i++ {
		node := &ListNode{Val: i}
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

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 每遍历一个节点，存储上一个遍历节点，再将当前节点指向上一个节点
 * 暂存原链表当前节点的后一个节点，再将该节点指向当前节点
 * 节点只遍历一遍，时间复杂度 O(N)，空间复杂度 O(1)
 */
func reverseList_prev(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

/*
 * 用递归的方式翻转链表
 * 实际上就是先遍历链表的每一个节点，没遍历一遍都存储当前节点
 * 到了最后一个节点，开始将该节点的 next 指针指向上一个节点，再断开上一个节点的指针
 * 此时上上个节点依然指向翻转后部份的尾节点，重复上述过程，就会从后往前依次翻转链表
 * 时间复杂度依然为 O(N)，空间复杂度也为 O(N)，需要存储 n 个节点，且每一次递归都需要存储占空间
 * 链表过长可能会导致栈溢出错误
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
