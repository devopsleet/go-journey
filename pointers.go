package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

func main() {

	Node1 := &ListNode{val: 5}
	Node2 := &ListNode{val: 10}
	Node3 := &ListNode{val: 15}

	Node1.Next = Node2
	Node2.Next = Node3

	head := Node1

	for head != nil {

		val := head.val

		fmt.Println("The value is  ", val)

		head = head.Next
	}

}
