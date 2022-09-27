package piscine

import (
	"fmt"
)

type NodeL struct {
	Data string
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data string) {
	a := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = a
		l.Tail = a
	}
	l.Tail.Next = a
	l.Tail = l.Tail.Next
}

func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
