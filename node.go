package linkedlist

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) toString() string {
	return fmt.Sprint(n.data)
}

func newNode(d interface{}) *Node {
	return &Node{
		data: d,
		next: nil,
	}
}
