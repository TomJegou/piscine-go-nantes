package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root != nil {
		current_node := root
		for current_node.Data != node.Data {
			if node.Data < current_node.Data {
				if current_node.Left != nil {
					current_node = current_node.Left
				}
			} else {
				if current_node.Right != nil {
					current_node = current_node.Right
				}
			}
		}
		if current_node.Parent != nil {

		}
	}
	return root
}
