package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 先遍历链表到数组中，再使用双指针判断数组是否为回文即可
 */
func isPalindrome_array(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodes := make([]int, 0)
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	start, end := 0, len(nodes)-1
	for start <= end {
		if nodes[start] != nodes[end] {
			return false
		}
		start++
		end--
	}
	return true
}

/*
 * 通过快慢指针找到链表中间节点，再翻转后半链表
 * 对比翻转后的链表以及长度和翻转前的链表
 * 返回结果前最好恢复链表的顺序
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	middleNode := middleNode(head)
	first, second := reverse(middleNode), head
	// 恢复链表
	defer reverse(middleNode)
	for first != nil {
		if first.Val != second.Val {
			return false
		}
		first, second = first.Next, second.Next
	}
	return true
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
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
