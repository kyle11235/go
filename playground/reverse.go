package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func create(count int) *Node {
	var head *Node
	var current *Node
	for i := 0; i < count; i++ {
		node := &Node{
			i,
			nil,
		}
		if head == nil {
			head = node
			current = head
		} else {
			current.Next = node
			current = node
			fmt.Printf("current=%v\n", current.Value)
		}
	}
	return head
}

func print(head *Node) {
	for head != nil {
		fmt.Printf("%v, ", head.Value)
		head = head.Next
	}
	println("")
}

// 123 -> 321
func reverse(head *Node) *Node {
	current := head

	// move
	for current != nil && current.Next != nil && current.Next.Next != nil {
		current = current.Next
	}

	// check quit
	if current.Next == nil {
		return head
	}

	// break
	newHead := current.Next
	current.Next = nil

	// deep
	deep := reverse(head)

	// reconnect
	newHead.Next = deep

	return newHead
}

// 012 345 678 9 -> 210 543 876 9
func reverseGroup(head *Node, k int) *Node {
	current := head

	// move
	for i := 0; i < k-1 && current != nil; i++ {
		current = current.Next
	}

	// check quit
	if current == nil {
		return head
	}

	// break
	toDeep := current.Next
	current.Next = nil
	newHead := reverse(head)

	// deep
	deep := reverseGroup(toDeep, k)

	// reconnect
	head.Next = deep

	return newHead
}

func main() {
	head := create(10)

	print(head)
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9,

	print(reverseGroup(head, 3))
	// 2, 1, 0, 5, 4, 3, 8, 7, 6, 9,

}
