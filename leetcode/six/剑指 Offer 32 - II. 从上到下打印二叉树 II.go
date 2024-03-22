package six

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//剑指 Offer 32 - III. 从上到下打印二叉树 III
func levelOrder2(root *TreeNode) [][]int {
	// 这tmd是声明切片但是这个切片里面每一个元素是[]int类型
	//ant := make([][]int, 0)
	//if root == nil {
	//	return ant
	//}
	//queue := list.New()
	//queue.PushBack(root)
	//ant = append(ant, []int{root.Val})
	//for queue.Len() != 0 {
	//	line := make([]int, 0)
	//	tem := queue.Front().Value.(*TreeNode)
	//	queue.Remove(queue.Front())
	//	if tem.Left != nil {
	//		queue.PushBack(tem.Left)
	//		line = append(line, tem.Left.Val)
	//	}
	//	if tem.Right != nil {
	//		queue.PushBack(tem.Right)
	//		line = append(line, tem.Right.Val)
	//	}
	//	if len(line) != 0 {
	//		ant = append(ant, line)
	//	}
	//
	//}
	//return ant
	//用二维数组遍历保存每一层的元素的值
	//思路如下：1.建立了两个结构体切片p、q,q相当于中间量，来确定每一行的元素个数，进行遍历保存在二维切片中,p是用来保存每一行的节点数
	//添加了一部分内容使其左右反转输出，存的时候反转一下就好，上一次的翻转使下一次的顺从改变，那么下次我们就要再翻转回来
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; 0 < len(q); i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		if i%2 == 0 {
			for j := 0; j < len(q); j++ {
				node := q[j]
				ret[i] = append(ret[i], node.Val)
				if node.Left != nil {
					p = append(p, node.Left)
				}
				if node.Right != nil {
					p = append(p, node.Right)
				}

			}
			q = rev(p)
		} else {
			for j := 0; j < len(q); j++ {
				node := q[j]
				ret[i] = append(ret[i], node.Val)
				if node.Right != nil {
					p = append(p, node.Right)
				}
				if node.Left != nil {
					p = append(p, node.Left)
				}
			}
			q = rev(p)
		}
	}
	return ret
}
func rev(slice []*TreeNode) []*TreeNode {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
