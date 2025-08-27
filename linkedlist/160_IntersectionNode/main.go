package main

func main() {

}

/*
 * 使用双指针，A 指针遍历 A 链表后遍历 B 链表
 * B 指针遍历 B 链表后遍历 A 链表
 * 无论是否相交，最多各自遍历 M+N 个节点后都会出结果
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

/*
 * 先存储链表 A 各个节点到 map 中，再遍历链表 B 的节点是否在 map 中
 * 相同的节点则表示链表相交，指针类型对象指向内存地址相同，则值一定相同
 * 因为都只遍历一遍，时间复杂度为 O(M+N)，空间复杂度 O(M)
 */
func getIntersectionNode_map(headA, headB *ListNode) *ListNode {
	headAPos := make(map[*ListNode]bool)
	for headA != nil {
		headAPos[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if headAPos[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
