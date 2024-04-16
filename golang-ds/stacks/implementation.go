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
