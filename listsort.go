package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func find_smallest_node(l *NodeI) *NodeI {
	min := l
	current_node := l
	for current_node != nil {
		if current_node.Data < min.Data {
			min = current_node
		}
		current_node = current_node.Next
	}
	return min
}

func ListSort(l *NodeI) *NodeI {
	current_node := l
	for current_node != nil {
		min := find_smallest_node(current_node)
		current_node.Data, min.Data = min.Data, current_node.Data
		current_node = current_node.Next
	}
	return l
}
