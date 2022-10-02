package piscine

func apply_func_lvl(root *TreeNode, levl int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if levl == 1 {
		f(root.Data)
	} else if levl > 1 {
		apply_func_lvl(root.Left, levl-1, f)
		apply_func_lvl(root.Right, levl-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for n := 1; n < BTreeLevelCount(root)+1; n++ {
		apply_func_lvl(root, n, f)
	}
}
