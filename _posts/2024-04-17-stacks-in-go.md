---
title: "Stacks in Golang"
date: 2024-04-17
categories: [Stacks]
tags: [Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Stacks]
---


# Stacks in Golang

[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/stacks-in-go.svg)](https://hits.sh/mokhlesurr031.github.io/posts/stacks-in-go/)


In Golang, a stack is a linear data structure that works on the Last-In-First-Out (LIFO) principle which means that the element which is pushed in the stack in last will be popped out first. 

In general stacks has 2 operations_

1. **Push:** Adds an element to the top of the stack.
2. **Pop:** Removes and returns the element from the top of the stack.

But we can also implement_

3. **Peek (or Top):** Returns the element from the top of the stack without removing it.
4. **isEmpty:** Checks if the stack is empty.
5. **Size (or Length):** Returns the number of elements currently in the stack.


<div style="text-align:center">
    <img src="static/stacks.png" alt="Linked List">
    <p>Fig: Stacks</p>
</div>


Implementation of Stack in Golang_

```go
package main

import "fmt"

type stack struct {
	items []interface{}
}

func (s *stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) Print() {
	fmt.Println("Stack:")
	for _, item := range s.items {
		fmt.Printf("%d->", item)
	}
	fmt.Println("\n")
}

func (s *stack) Length() int {
	return len(s.items)
}

func main() {
	stack := stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()

	ln := stack.Length()
	fmt.Println("Length:", ln)
	fmt.Println("\n")

	fmt.Println("Peek:", stack.Peek())

	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())

	fmt.Println("Is empty?", stack.IsEmpty())
}
```

Output 

```bash
Stack:
1->2->3->

Length: 3


Peek: 3
Pop: 3
Pop: 2
Pop: 1
Is empty? true
```





