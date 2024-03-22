package seven

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return recur(root.Left, root.Right)
//}
//
//func recur(L, R *TreeNode) bool {
//	if L == nil && R == nil {
//		return true
//	}
//	if (L == nil && R != nil) || (L != nil && R == nil) {
//		return false
//	}
//	if L.Val != R.Val {
//		return false
//	}
//	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
//}
