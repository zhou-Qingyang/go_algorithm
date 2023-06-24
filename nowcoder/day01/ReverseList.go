package day01

type ListNode struct {
	Val  int
	Next *ListNode
}

//BM1 反转链表
func ReverseList(head *ListNode) *ListNode {
	// write code here
	pre := &ListNode{}
	cur := head
	for cur != nil {
		currentNode := cur.Next //获取 下一个节点
		cur.Next = pre          //下一个节点 变换到最前面
		pre = cur               //移动指针
		cur = currentNode       //移动指针

	}
	return pre

}
