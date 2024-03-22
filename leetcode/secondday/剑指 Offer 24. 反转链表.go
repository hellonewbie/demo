package secondday

type ListNode struct {
	Val  int
	Next *ListNode
}

//改变链表的指向
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
