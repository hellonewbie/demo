package secondday

//借助辅助栈进行链表的逆序输出

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reversePrint(head *ListNode) []int {
//	var (
//		list  []int
//		stack []int
//		top   *int
//	)
//	for head != nil {
//		stack = append(stack, head.Val)
//		head = head.Next
//	}
//	for len(stack) >= 1 {
//		top = &stack[len(stack)-1]
//		list = append(list, *top)
//		stack = stack[:len(stack)-1]
//	}
//	return list
//}
