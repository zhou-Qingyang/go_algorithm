package day01

type ListNode struct {
	Val  int
	Next *ListNode
}

//19 删除链表的倒数第 N 个结点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	p, q := dummy, dummy

	for i := 0; i < n; i++ {
		q = q.Next
	}

	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
