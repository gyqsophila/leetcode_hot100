package main

func main() {

}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 还是将遍历的节点放进 map，后续有相同节点则遇到了环，返回当前节点即可
 */
func detectCycle_map(head *ListNode) *ListNode {
	tmp := make(map[*ListNode]bool)
	for head != nil {
		if tmp[head] {
			return head
		}
		tmp[head] = true
		head = head.Next
	}
	return nil
}

/*
 * 快慢指针+第二个慢指针，快慢指针第一次相遇后，第二个慢指针从 head 出发，当 2 个慢指针第一次相遇的时候就是进入环的节点
 * 快慢第一次相遇后，快指针的路程(y)是慢指针路程(x)的 2 倍，此时第二个慢指针出发，走过 x 的路程，到达快慢第一次相遇点，此时第一个慢指针也回到了该节点(2x=y)
 * 也就是说 2 个慢指针必然会相遇，那么第一次相遇的点必然是入环的节点
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		// 第一次相遇
		if slow == fast {
			slow2 := head
			for slow2 != slow {
				slow2 = slow2.Next
				slow = slow.Next
			}
			return slow2
		}
	}
	return nil
}
