package day03

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	sum := 0
	carry := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		sum = l1.Val + l2.Val + carry
		carry = sum % 10

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next

		// 继续遍历 l1 和 l2
		l1 = l1.Next
		l2 = l2.Next

	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return cur.Next
}
