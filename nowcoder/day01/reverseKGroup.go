package day01

type ListNode struct {
	Val  int
	Next *ListNode
}

//BM3 链表中的节点每k个一组翻转
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 创建一个哑节点作为虚拟头节点
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head
	count := 0

	// 计算链表的长度
	for curr != nil {
		count++
		curr = curr.Next
	}

	// 依次处理每一组节点
	for count >= k {
		curr = prev.Next
		next := curr.Next
		for i := 1; i < k; i++ {
			// 将当前节点移至下一组的开头位置
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
			next = curr.Next
		}
		prev = curr
		count -= k
	}

	return dummy.Next
}
