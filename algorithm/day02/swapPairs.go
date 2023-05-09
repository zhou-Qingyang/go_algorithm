package day02

type ListNode struct {
	Val  int
	Next *ListNode
}

//24
func SwapPairs(head *ListNode) *ListNode {
	dummt := &ListNode{0, head}
	prev := dummt

	for prev.Next != nil && prev.Next.Next != nil {
		node1 := prev.Next
		node2 := node1.Next
		nextNode := node2.Next

		node2.Next = node1
		node1.Next = nextNode
		prev.Next = node2
	}

	return dummt.Next
}
