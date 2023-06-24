package day01

type ListNode struct {
	Val  int
	Next *ListNode
}

//BM4 合并两个排序的链表
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	l1 := pHead1
	l2 := pHead2

	dummy := &ListNode{Val: -1, Next: nil}
	for l1 != nil || l2 != nil {
		if l1.Val > l2.Val {
			dummy.Next = l2
			l2 = l2.Next
		} else {
			dummy.Next = l1
			l1 = l1.Next
		}
		dummy = dummy.Next
	}
	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}

	return dummy.Next

}
