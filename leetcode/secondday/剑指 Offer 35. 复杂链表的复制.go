package secondday

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//链表复制完成(只是简答的值复制)
//func copyRandomList(head *Node) *Node {
//	var ListNode *Node
//	NodeHeader := new(Node)
//	ListNode = NodeHeader
//	NodeHeader.Val = head.Val
//	for head.Next != nil {
//		head = head.Next
//		NodeE := new(Node)
//		ListNode.Next = NodeE
//		NodeE.Val = head.Val
//		ListNode = NodeE
//	}
//	return NodeHeader
//}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//执行之后对变量进行更新
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
