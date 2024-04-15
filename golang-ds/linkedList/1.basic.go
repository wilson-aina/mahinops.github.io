package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) insert(data int) {
	newNode := &node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode
}

func (l *linkedList) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println("\n")
}

func main() {
	list := linkedList{}
	fmt.Println(list)
	list.insert(3)
	list.insert(4)

	fmt.Println("Linked List:")
	list.print()

}
