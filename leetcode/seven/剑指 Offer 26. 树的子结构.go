package seven

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//总结一下弱小的我花了1h做出来的题目吧
//卡我的点有:我该怎么遍历,思维定势取值比较，问题就是你怎么知道这个结构是你想的那样万一子
//结构和同层的值相同怎么办，这个时候我们就要借用一个辅助函数了，遍历到哪个点的值与子结构
//根节点一样的值，就用以哪个点为根节点在辅助函数里面进行比对
//递归的思路:1.结束条件(寻常条件、特殊条件)
//这个也卡了我一会儿，我在里面建立了一个条件判断，只要一个不满足就直接返回了
//存在更深层的树满足条件的存在，然后我们就可以加入递归条件中去解决这个问题
//2.递归的参数选择
//函数的执行顺序也很关键
//func isSubStructure(A *TreeNode, B *TreeNode) bool {
//	if A == nil || B == nil {
//		return false
//	}
//
//	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
//}
//
//func recur(A *TreeNode, B *TreeNode) bool {
//	if B == nil {
//		return true
//	}
//	if A == nil {
//		return false
//	}
//
//	if A.Val != B.Val {
//		return false
//	}
//	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
//}
