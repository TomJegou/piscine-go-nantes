package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	result := 1
	var (
		treelevl_left  int
		treelevl_right int
	)
	if root.Left != nil {
		treelevl_left = result + BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		treelevl_right = result + BTreeLevelCount(root.Right)
	}
	if treelevl_left > treelevl_right {
		return treelevl_left
	}
	return treelevl_right
}
