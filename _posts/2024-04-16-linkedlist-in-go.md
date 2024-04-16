---
title: "Linked List in Golang"
date: 2024-04-16
categories: [LinkedList]
tags: [Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Linked List]
---


# Linked List in Golang
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/linkedlist-in-go.svg)](https://hits.sh/mokhlesurr031.github.io/posts/linkedlist-in-go/)


Linked List is a linear data structure like arrays, whereas, linked list doesn't have element stored in contiguous locations, as arrays do. 

In simple words, we can say, linked list is a collection of `Nodes`. 

`Nodes` consists of two parts_

1. **Data:** The actual value stored in the node.
2. **Pointer:** A reference to the next node in the sequence.

<div style="text-align:center">
    <img src="static/linked_list.png" alt="Linked List">
    <p>Fig: Linked List</p>
</div>

Each `node` in the linked list contains the `data` part and the `pointer` to the next node in the linked list.

The linked list supported operations are_
- Insertion
  - Insert at the beginning
  - Insert at the end
  - Insert at a given position

- Deletion
  - Delete from the beginning
  - Delete from the end
  - Delete at a given position
  - Delete a node with a given value

- Traversal
  - Traverse the list and print values

- Searching
  - Search for a given value

- Length
  - Find the length of the list

- Reversal
  - Reverse the linked list

- Sorting
  - Sort the linked list

- Concatenation
  - Concatenate two linked lists



In Go, you can define a linked list node as a struct_

```go
// Node represents a node in the linked list
type node struct {
	data int
	next *node
}
```

Defining a linked list using the node is like_
```go
// LinkedList represents a linked list
type linkedList struct {
	head *node
}
```

### Lets complete this part_

```go
package main

import "fmt"

// defining node with data and address to the next node
type node struct {
	data int
	next *node
}

// define linked list head with the node
type linkedList struct {
	head *node
}

func main() {
  // initiate linked list currently empty
	list := linkedList{}

  // lets write and call a function to travers and print all the elements of linked list
  list.print()

  // cool, we have written the print function that will travers and print all linked list elements
  // as we have already called the function, it will return nothing. 

  // lets write a function to insert element in linked list at last index
  list.insert(1)
  list.insert(2)
  list.insert(3)

}

func (l *linkedList) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println("\n")
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

```


Now lets complete the linked list operations one by one_


```go
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

```

Output

```bash
Empty List
{<nil>}


Traverse


Insert at the beginning 1
Insert at the beginning 2


Insert at the beginning 3


Insert at the beginning 4


Insert at the beginning 5


Traverse
5->4->3->2->1->

Insert at the end 7


Insert at the end 8


Insert at the end 9


Insert at the end 10


Insert at the end 11


Traverse
5->4->3->2->1->7->8->9->10->11->

Inserted 6 at position 6
Traverse
5->4->3->2->1->6->7->8->9->10->11->

Deleting from beginning


Traverse
4->3->2->1->6->7->8->9->10->11->

Deleting from end
Traverse
4->3->2->1->6->7->8->9->10->

Traverse
4->2->1->6->7->8->9->10->

Traverse
2->1->6->7->8->9->10->

Value 2 found at position 1
Found values 2 at position->1

Value 3 not found in the list.
Found values 3 at position->-1

Traverse
2->1->6->7->8->9->10->

Length of the list: 7
Traverse
2->1->6->7->8->9->10->

Traverse
10->9->8->7->6->1->2->

Traverse
1->2->6->7->8->9->10->

```




