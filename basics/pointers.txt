package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	dict := make(map[*ListNode]bool)

	ptr := head

	for ptr != nil {
		if _, ok := dict[ptr]; ok {
			return true
		}

		dict[ptr] = true
		ptr = ptr.Next
	}

	return false

}

func main() {

	Node1 := &ListNode{val: 5}
	Node2 := &ListNode{val: 10}
	Node3 := &ListNode{val: 15}

	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node1

	head := Node1

	output := hasCycle(head)
	fmt.Println(output)

}
