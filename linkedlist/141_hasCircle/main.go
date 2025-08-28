package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 将遍历过的节点放进 map，于后续遍历的节点进行比较，若相同则存在环
 * 时间复杂度 O(N)，空间复杂度 O(N)
 */
func hasCycle_map(head *ListNode) bool {
	tmp := make(map[*ListNode]bool)
	for head != nil {
		if tmp[head] {
			return true
		}
		tmp[head] = true
		head = head.Next
	}
	return false
}

/*
 * 快慢指针先后出发，快指针每次前进 2，慢指针每次前进 1，若两指针还能相遇，则说明有环
 * 时间复杂度 O(N)，空间复杂度 O(1)
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

/*
 * 快慢指针同时出发，快指针每次前进 2，慢指针每次前进 1，若两指针还能再次相遇，则说明有环
 * 时间复杂度 O(N)，空间复杂度 O(1)
 */
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
