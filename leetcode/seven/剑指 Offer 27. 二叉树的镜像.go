package seven

//卡住了
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func isSymmetric(root *TreeNode) bool {
	if mirrorTree(root) == nil {
		return false
	}
	root2 := mirrorTree(root)
	return Check(root, root2)

}

func Check(root *TreeNode, root2 *TreeNode) bool {
	if root == nil && root2 == nil {
		return true
	}
	if (root == nil && root2 != nil) || (root != nil && root2 == nil) {
		return false
	}
	if root.Val != root2.Val {
		return false
	}
	return Check(root.Left, root2.Left) && Check(root.Right, root2.Right)
}
