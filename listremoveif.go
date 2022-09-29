package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head.Data == data_ref && l.Head != nil {
		if l.Head.Next != nil {
			l.Head = l.Head.Next
		} else {
			l.Head = nil
			l.Tail = nil
			return
		}
	}
	nodel_after := l.Head
	nodel_prev := &NodeL{Next: nodel_after}
	for nodel_after != nil {
		if nodel_after == l.Tail && nodel_after.Data == data_ref {
			l.Tail = nodel_prev
			l.Tail.Next = nil
		} else {
			if nodel_after.Data == data_ref {
				nodel_prev.Next = nodel_after.Next
				nodel_after = nodel_after.Next
				continue
			}
		}
		nodel_prev = nodel_after
		nodel_after = nodel_after.Next
	}
}
