package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Data < root.Data {
			return BTreeIsBinary(root.Left)
		} else {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Data > root.Data {
			return BTreeIsBinary(root.Right)
		} else {
			return false
		}
	}
	return true
} //test
