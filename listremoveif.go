package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if ListSize(l) > 0 {
		if l.Head.Data == data_ref {
			if l.Head.Next != nil {
				l.Head = l.Head.Next
			} else {
				l.Head = nil
				l.Tail = nil
			}
		}
		current_node := l.Head
		for current_node != nil {
			if current_node.Next == l.Tail {
				if current_node.Next.Data == data_ref {
					l.Tail = current_node
					current_node.Next = nil
				}
			} else if current_node.Next.Data == data_ref {
				current_node.Next = current_node.Next.Next
				continue
			}
			current_node = current_node.Next
		}
	}
}
