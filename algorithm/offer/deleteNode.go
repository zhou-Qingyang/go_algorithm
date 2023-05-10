package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

//18
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 头结点就是要删除的节点
	if head.Val == val {
		head = head.Next
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head

}
