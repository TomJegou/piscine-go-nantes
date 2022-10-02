package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
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
		rplc.Left = current_node.Left
		rplc.Right = current_node.Right
		current_node.Data = rplc.Data
	}
	return root
}
