package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	for node.Data != root.Data {
		if root == nil {
			root = node
		}
		if node.Data < root.Data {
			root = root.Left
		} else if node.Data > root.Data {
			root = root.Left
		}
	}
	root.Data = rplc.Data
	return root
}
