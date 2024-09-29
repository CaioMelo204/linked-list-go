package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func traverse(head *Node) {
	cur := head
	fmt.Println("Taverse the LinkedList")
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}

func push(head *Node, data int) *Node {
	newNode := &Node{data: data}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return head
}

func pop(head *Node) *Node {
	current := head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	return head
}

func unshift(head *Node, data int) *Node {
	newNode := &Node{data: data}
	newNode.next = head

	head = newNode
	return newNode
}

func shift(head *Node) *Node {
	head = head.next
	return head
}

func findLength(head *Node) int {
	current := head
	i := 0
	for current != nil {
		current = current.next
		i++
	}
	return i
}

func main() {
	list := &Node{data: 1}
	list = push(list, 2)
	list = push(list, 3)
	traverse(list)
	list = unshift(list, 0)
	list = unshift(list, -1)
	list = unshift(list, -2)
	traverse(list)
	list = shift(list)
	traverse(list)
	list = shift(list)
	traverse(list)
	list = pop(list)
	traverse(list)
	list = pop(list)
	traverse(list)
	fmt.Println("length: ", findLength(list))
}
