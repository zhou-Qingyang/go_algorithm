package day01

type ListNode struct {
	Val  int
	Next *ListNode
}

//BM2 链表内指定区间反转
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if m == n || head == nil {
		return nil
	}
	dummy := &ListNode{Val: -1, Next: head}
	prev := dummy
	for i := 1; i < m; i++ {
		prev = prev.Next
	}
	current := prev.Next //现在指的就是第m个节点

	reverseNode := &ListNode{Val: -1, Next: nil}
	for i := m; i < n; i++ {
		currNode := current.Next
		current.Next = reverseNode
		reverseNode = current
		current = currNode
	}
	//这两行 不是很懂
	prev.Next.Next = current
	prev.Next = reverseNode
	//
	return dummy.Next
}
