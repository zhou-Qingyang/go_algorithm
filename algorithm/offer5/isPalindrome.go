func isPalindrome(head *ListNode) bool {
    // 如果链表为空或只有一个节点，直接返回 true
    if head == nil || head.Next == nil {
        return true
    }
    // 使用快慢指针找到链表中点
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 反转后半部分链表
    cur, prev := slow.Next, slow
    slow.Next = nil
    for cur != nil {
        temp := cur.Next
        cur.Next = prev
        prev, cur = cur, temp
    }
    // 比较链表前半部分和后半部分是否相同
    p1, p2 := head, prev
    for p2 != nil {
        if p1.Val != p2.Val {
            return false
        }
        p1, p2 = p1.Next, p2.Next
    }
    return true
}