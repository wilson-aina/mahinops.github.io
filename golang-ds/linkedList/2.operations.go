package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (ll *linkedList) traverse() {
	fmt.Println("Traverse")
	current := ll.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println("\n")
}

func (ll *linkedList) insertAtBeginning(data int) {
	fmt.Println("Insert at the beginning", data)
	newNode := &node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode

	fmt.Println("\n")
}

func (ll *linkedList) insertAtEnd(data int) {
	fmt.Println("Insert at the end", data)
	newNode := &node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	lastNode := ll.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode
	fmt.Println("\n")
}

func (ll *linkedList) insertAtPosition(position, data int) {
	if position < 1 {
		fmt.Println("Invalid position")
		return
	}

	if position == 1 {
		ll.insertAtBeginning(data)
		return
	}

	newNode := &node{data: data, next: nil}
	current := ll.head
	previous := ll.head
	count := 1

	for current != nil && count < position {
		previous = current
		current = current.next
		count++
	}

	if count == position {
		previous.next = newNode
		newNode.next = current
		fmt.Println("Inserted", data, "at position", position)
	} else {
		fmt.Println("Position out of range")
	}
}

func (ll *linkedList) deleteFromBeginning() {
	if ll.head == nil {
		fmt.Println("List is empty. Nothing to delete.")
		return
	}

	fmt.Println("Deleting from beginning")

	// Move the head to the next node
	ll.head = ll.head.next
	fmt.Println("\n")

}

func (ll *linkedList) deleteFromEnd() {
	if ll.head == nil {
		fmt.Println("List is empty. Nothing to delete.")
		return
	}
	fmt.Println("Deleting from end")
	// If there's only one node in the list
	if ll.head.next == nil {
		ll.head = nil
		return
	}

	// Find the second to last node
	secondToLast := ll.head
	for secondToLast.next.next != nil {
		secondToLast = secondToLast.next
	}
	// Set the next of second to last node to nil
	secondToLast.next = nil

}

func (ll *linkedList) deleteFromPosition(position int) {
	if position < 1 {
		fmt.Println("Invalid position")
		return
	}

	if ll.head == nil {
		fmt.Println("List is empty. Nothing to delete.")
		return
	}
	if position == 1 {
		ll.deleteFromBeginning()
		return
	}

	current := ll.head
	previous := ll.head
	count := 1

	for current != nil && count < position {
		previous = current
		current = current.next
		count++
	}

	if current == nil {
		fmt.Println("Position out of range")
		return
	}

	previous.next = current.next

}

func (ll *linkedList) deleteByValue(value int) {
	if ll.head == nil {
		fmt.Println("List is empty. Nothing to delete.")
		return
	}

	// If the value to delete is in the head node
	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	previous := ll.head

	for current != nil && current.data != value {
		previous = current
		current = current.next
	}

	if current == nil {
		fmt.Println("Value not found in the list.")
		return
	}

	previous.next = current.next

}

func (ll *linkedList) searchByValue(value int) int {
	if ll.head == nil {
		fmt.Println("List is empty. Value not found.")
		return -1
	}

	current := ll.head
	position := 1

	for current != nil {
		if current.data == value {
			fmt.Println("Value", value, "found at position", position)
			return position
		}
		current = current.next
		position++
	}

	fmt.Println("Value", value, "not found in the list.")
	return -1

}

func (ll *linkedList) findLength() {
	length := 0
	current := ll.head

	for current != nil {
		length++
		current = current.next
	}

	fmt.Println("Length of the list:", length)

}

func (ll *linkedList) reverse() {
	var prev, next *node
	current := ll.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	ll.head = prev

}

func (ll *linkedList) sorting() {
	if ll.head == nil || ll.head.next == nil {
		// List has 0 or 1 element, no need to sort
		return
	}

	// Bubble sort
	sorted := false
	for !sorted {
		sorted = true
		prev := ll.head
		current := ll.head.next

		for current != nil {
			if prev.data > current.data {
				// Swap data
				prev.data, current.data = current.data, prev.data
				sorted = false
			}
			prev = current
			current = current.next
		}
	}
}

// func (ll *linkedList) concate(other *linkedList) {
// 	if ll.head == nil {
// 		// If the first list is empty, simply set its head to the head of the second list
// 		ll.head = other.head
// 		return
// 	}

// 	if other.head == nil {
// 		// If the second list is empty, nothing to concatenate
// 		return
// 	}

// 	// Traverse to the end of the first list
// 	current := ll.head
// 	for current.next != nil {
// 		current = current.next
// 	}

// 	// Concatenate the second list to the end of the first list
// 	current.next = other.head
// }

func main() {
	list := linkedList{}

	fmt.Println("Empty List")
	fmt.Println(list)
	fmt.Println("\n")

	list.traverse()

	list.insertAtBeginning(1)
	list.insertAtBeginning(2)
	list.insertAtBeginning(3)
	list.insertAtBeginning(4)
	list.insertAtBeginning(5)

	list.traverse()

	list.insertAtEnd(7)
	list.insertAtEnd(8)
	list.insertAtEnd(9)
	list.insertAtEnd(10)
	list.insertAtEnd(11)

	list.traverse()

	list.insertAtPosition(6, 6)

	list.traverse()

	list.deleteFromBeginning()

	list.traverse()

	list.deleteFromEnd()

	list.traverse()

	list.deleteFromPosition(2)

	list.traverse()

	list.deleteByValue(4)

	list.traverse()

	position := list.searchByValue(2)
	fmt.Printf("Found values 2 at position->%d", position)
	fmt.Println("\n")

	position1 := list.searchByValue(3)
	fmt.Printf("Found values 3 at position->%d", position1)
	fmt.Println("\n")

	list.traverse()

	list.findLength()

	list.traverse()

	list.reverse()

	list.traverse()

	list.sorting()

	list.traverse()

}
