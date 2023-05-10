package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

//06
func reversePrint(head *ListNode) []int {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	n := len(stack)
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, stack[n-i-1])
	}
	return arr

}
