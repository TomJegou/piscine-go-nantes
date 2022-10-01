package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left.Parent = root
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right.Parent = root
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}
