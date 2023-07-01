package linkedlist

// package main

// import (
// 	"fmt"

// 	linkedlist "github.com/nqnlong1506/linked-list/linked-list"
// )

// type Product struct {
// 	name  string
// 	price int
// }

// func main() {
// 	list := linkedlist.NewLinkedList()
// 	list.InsertToEnd("nqnlong")
// 	list.InsertToEnd(12)
// 	list.InsertToHead("check checks")
// 	list.InsertToEnd("nguyen bay")
// 	list.InsertToHead("first commit")
// 	list.Print()

// 	fmt.Println()
// 	fmt.Println(list.LastNode())

// 	node := list.FindNode("nguyen bay")

// 	fmt.Println("reult", node)
// 	list.AddAfterNode("nqnlong", "ngoclong")
// 	l := list.InterateList()
// 	for _, v := range l {
// 		fmt.Println(v)
// 	}

// 	node = list.FindNode("ngoclong")
// 	fmt.Println(*node)

// 	list.Print()
// 	fmt.Println("delete node")

// 	err := list.DeleteNode("ngoclong")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		list.Print()
// 	}

// 	node = list.FindNode("ngoclong")
// 	if node == nil {
// 		fmt.Println("does not exist")
// 	} else {
// 		fmt.Println(*node)
// 	}

// }

/*
	struct node
	struct linkedlist

	addtohead // insert to head
	interateList
	lastNode // get last node of linked list
	addtoend // insert to end
	node with value // search node with specified data
	addafter // insert after given node
*/
