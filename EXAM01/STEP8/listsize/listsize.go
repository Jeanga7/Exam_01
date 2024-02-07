package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head}
	l.Head = newNode
	if newNode.Next == nil {
		l.Tail = newNode
	}
}

func ListSize(l *List) int {
	size := 0

	for l.Head != nil {
		size++
		l.Head = l.Head.Next
	}
	return size
}


func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}
