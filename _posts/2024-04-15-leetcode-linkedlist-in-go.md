---
title: "Linked List in Golang"
date: 2024-04-15
categories: [LinkedList]
tags: [Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Linked List]
---


# Linked List in Golang

Linked List is a linear data structure like arrays, whereas, linked list doesn't have element stored in contiguous locations, as arrays do. 

In simple words, we can say, linked list is a collection of `Nodes`. 

`Nodes` consists of two parts_

1. **Data:** The actual value stored in the node.
2. **Pointer:** A reference to the next node in the sequence.

Each `node` in the linked list contains the `data` part and the `pointer` to the next node in the linked list.

The linked list supported operations are_
- Insertion
  - Insert at the beginning
  - [Insert at the end](### Lets complete this part_)
  - Insert at a given position

- Deletion
  - Delete from the beginning
  - Delete from the end
  - Delete at a given position
  - Delete a node with a given value

- Traversal
  - [Traverse the list and print values](### Lets complete this part_)

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

```
// Node represents a node in the linked list
type node struct {
	data int
	next *node
}
```

Defining a linked list using the node is like_
```
// LinkedList represents a linked list
type linkedList struct {
	head *node
}
```

### Lets complete this part_

```
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








