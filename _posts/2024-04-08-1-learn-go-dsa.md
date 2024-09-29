---
title: "Learn Go with Data Structure and Algorithm-Part-1"
date: 2024-04-08
categories: [Golang]
tags: [Go, Golang, DSA, Problem Solving, Data Structure, Algorithm, Programming Language]
---

# 01-Go Notebook
[![Hits](https://hits.sh/mahinops.github.io/posts/1-learn-go-dsa.svg)](https://hits.sh/mahinops.github.io/posts/1-learn-go-dsa/)


## Data Structure

### Concept of Stack 

A stack is a memory in which values are stored and retrieved in "Last in First Out" manner. 

Data is added to the stack using `Push` operation and data is taken out of the stack using `Pop` operation.


### Recursive Function
A recursive funtion is a function that calls itself, directly or indirectly. 

A recursive function consists of two parts_
1. Termination Condition.
2. Body. 

`Termination Condition:` A recursive method always contains one or more terminating condition. A condition in which recursive method process a simple case and do not call itself.

`Body: (including recursive expansion)` The main logic of the recursive method contained in the body of the method. It also contains the recursion expansion statement that in turn calls the method itself. 

Three important properties of recursive algorithm are:_
1. A recursive algorithm must have a termination condition. 
2. A recursive algorithm must change its state, and move towards the termination condition. 
3. A recursive algorithm must call itself. 

`Note:` The speed of a recursive program is slower because of stack overheads. If the same task can be done using an iterative solution (using loops), then we should prefer an iterative solution in place of recursion to avoid stack overhead.

`Note:` Without termination condition, the recursive method may run forever and will finally consume all the stack memory.


#### Example Problems

- Factorial calculation. `n! = n*(n-1)*(n-2)*...2*1`

```go
package main

import "fmt"

func main() {
	fmt.Println("factorial 5 is :: ", Factorial(5))
}

func Factorial(num int) int {
	// Termination Condition
	if num <= 1 {
		return 1
	}
	// Body, Recursive Expansion
	return num * Factorial(num-1)
}
```

- Find greatest common divisor.





