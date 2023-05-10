package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode = nil, nil
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			if head == nil {
				head = l1
				tail = l1
			} else {
				tail.Next = l1
				tail = tail.Next
			}
			l1 = l1.Next
		} else {
			if head == nil {
				head = l2
				tail = l2
			} else {
				tail.Next = l2
				tail = tail.Next
			}
			l2 = l2.Next
		}
	}
	if l1 != nil {
		if head == nil {
			return l1
		} else {
			tail.Next = l1
		}
	}
	if l2 != nil {
		if head == nil {
			return l2
		} else {
			tail.Next = l2
		}
	}
	return head

}
