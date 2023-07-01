package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		// fmt.Println(current.data)
		fmt.Println(current.toString())
		current = current.next
	}
}

// insert node to head
func (ll *LinkedList) InsertToHead(data interface{}) {
	newNode := newNode(data)

	// check exist head of linked list
	if ll.head == nil {
		ll.head = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode
}

// iterate list
func (ll *LinkedList) InterateList() []interface{} {
	var list []interface{}
	current := ll.head
	for {
		if current == nil {
			break
		}
		list = append(list, current.data)
		current = current.next
	}

	return list
}

// last node
func (ll *LinkedList) LastNode() *Node {
	current := ll.head
	if current == nil {
		return nil
	}
	for {
		if current.next == nil {
			return current
		}
		current = current.next
	}
}

// insert to end
func (ll *LinkedList) InsertToEnd(data interface{}) {
	// newNode := &Node{
	// 	data: data,
	// 	next: nil,
	// }
	newNode := newNode(data)

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// node with data
func (ll *LinkedList) FindNode(data interface{}) *Node {
	current := ll.head

	if current == nil {
		return nil
	}

	for {
		if current.data == data {
			return current
		}
		if current.next == nil {
			return nil
		}
		current = current.next
	}
}

// add after given data
func (ll *LinkedList) AddAfterNode(nodeData, data interface{}) {
	newNode := newNode(data)
	node := ll.FindNode(nodeData)
	if node == nil {
		if ll.head == nil {
			ll.head = newNode
			return
		} else {
			ll.InsertToEnd(data)
			return
		}
	}
	newNode.next = node.next
	node.next = newNode
}

func (ll *LinkedList) DeleteNode(data interface{}) error {
	current := ll.head

	if current == nil {
		return fmt.Errorf("does not contain node")
	}

	if current.data == data {
		ll.head = current.next
		return nil
	}

	for {
		if current.next == nil {
			return fmt.Errorf("does not contain node")
		}

		if current.next.data == data {
			// delete node
			current.next = current.next.next
			return nil
		}

		current = current.next
	}
}
