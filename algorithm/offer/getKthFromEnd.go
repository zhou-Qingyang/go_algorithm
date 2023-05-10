package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	for cur != nil {
		head = head.Next
		cur = cur.Next
	}
	return head

}
