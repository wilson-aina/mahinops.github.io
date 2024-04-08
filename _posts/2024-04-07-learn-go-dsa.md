---
title: "Learn Go with Data Structure and Algorithm-Part-0"
date: 2024-04-07
categories: [Golang]
tags: [Go, Golang, DSA, Problem Solving, Data Structure, Algorithm, Programming Language]
---

# 00-Go Notebook
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/learn-go-dsa.svg)](https://hits.sh/mokhlesurr031.github.io/posts/learn-go-dsa/)

## Basic Data Types

There are 8 important premitive data types
- Boolean
- Integer
- Unsigned Integer - Special types of integer, which can store only positive values. 
- Float
- Strings - Sequence of Unicode charaters, immutable.
- Byte - are alias of uint8.
- Rune - alias of int32 and used to store Unicode characters.
- Complex - store complex numbers. 


### Strings
Strings are immutable so you cannot change its content once created. You need to first convert into a slice of rune then do the change and in the end convert it back to string. 

Example
```
package main

import "fmt"

func main() {
	sampleStr := "Hello World!"
	fmt.Println(sampleStr)
	fmt.Println(sampleStr[0]) //output is 72

	// update
	r := []rune(sampleStr)
	r[0] = 'M'
	sampleStr = string(r)
	fmt.Println(sampleStr)
}
```


#### String Operations

- Length - len(stringVar)
- Concatenation - str1+str2
- Equality - "Hello"=="hello"
- Unicode Comparison - "a">"b"
- Slicing - strVar[1:4]


### Conditions and Loops

#### Basic if
```
if x>y{
	fmt.Println("X")
}
```

#### Precondition if
```
if area:=length*width; area<limit{
	fmt.Println("something"
}
```

#### For Loop
For loop helps to iterte through a group of statements mulitple times. 


Go has 4 forms of For Loop. 

##### Type 1
```
for <initialization>; <condition>; <increment/decrement> {}
---

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum is :: ", sum)
}
```
##### Type 2 - like a while loop
```
for <condition>{}
---

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	n := len(numbers)
	for i < n {
		sum += numbers[i]
		i++
	}
	fmt.Println("Sum is :: ", sum)
}
```

##### Type 3 - an infinite while loop
```
for {}
---

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	n := len(numbers)
	for {
		sum += numbers[i]
		i++
		if i >= n {
			break
		}
	}
	fmt.Println("Sum is :: ", sum)
}
```

##### Type 4 - with range
```
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum is :: ", sum)
}
```

#### Range
The range keyword is used in for loop to iterate data in data structure (arrays, slices, string, maps etc).

- Array/Slice - range provides indexes and values. 
- Maps - range provides key-value pairs. 
- String - range provides index of each Unicode characters in string and their corresponding Unicode characters. 

If index values are not needed then they can be discarded using underscore (_). 

Example:

```
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Print("[", index, ",", val, "] ")
	}
	fmt.Println("\nSum is :: ", sum)
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}
	str := "Hello, World!"
	for index, c := range str {
		fmt.Print("[", index, ",", string(c), "] ")
	}
}
```

Output
```
[0,1] [1,2] [2,3] [3,4] [4,5] [5,6] [6,7] [7,8] [8,9] [9,10] 
Sum is ::  55
1  ->  apple
2  ->  banana
[0,H] [1,e] [2,l] [3,l] [4,o] [5,,] [6, ] [7,W] [8,o] [9,r] [10,l] [11,d] [12,!] 
```


### Function 
Functions are used to provide modularity to the program. We can divide complex tasks into managable tasks using functions. 

Example
```
package main

import "fmt"

func main() {
	result := getMinMax(20, 40)
	fmt.Println(result)

}

func getMinMax(x, y int) string {
	if x > y {
		return fmt.Sprintf("%d is min and %d is max", y, x)
	} else {
		return fmt.Sprintf("%d is min and %d is max", x, y)
	}
}
```

Output
```
Value of i before increment is :  10
Value of i after increment is :  10
```

Analysis
- Variable “i” is declared and value 10 is initialized to it.
- Value of ”i” is printed.
- Increment function is called. When a function is called the value of the
parameter is copied into another variable of the called function. Flow of
control goes to line no 1.
- Value of var is incremented by 1. However, remember, it is just a copy
inside the increment function.
- When the function exits, the value of ”i” is still 10.


Points to remember:
1. Pass by value just creates a copy of variable.
2. Pass by value, value before and after the function call remain same.


### Pointer
Pointers are nothing more than variables that store memory address of another variable and can be used to access the value stored at those addresses. 

Example
```
package main

import "fmt"

func main() {
	data := 10
	ptr := &data
	fmt.Println("Value stored at variable var is ", data)
	fmt.Println("Value stored at variable var is ", *ptr)
	fmt.Println("The address of variable var is ", &data)
	fmt.Println("The address of variable var is ", ptr)
}
```

Output
```
Value stored at variable var is  10
Value stored at variable var is  10
The address of variable var is  0xc00009e018
The address of variable var is  0xc00009e018
```

Analysis
- An integer type variable var is created, which store value 10
- A pointer ptr is created, which store address of variable var.
- Value stored in variable var is printed to screen. Using * Operator value
stored at the pointer location is accessed.
- Memory address of var is printed to the screen. & operator is used to get
the address of a variable
	

#### Parameter Passing, Call by Pointer/Reference

If you need to change the value of the parameter inside the function, you should call by reference. C language by default passes the value. 

Therefore, to make it happen, you need to pass the address of a variable and changing the values of the variable using this address inside the called function. 

Example
```
package main

import "fmt"

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}
func main() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is : ", i)
}
```

Output
```
Value of i before increment is :  10
Value of i after increment is :  11
```

Analysis

- Address of “i” is passed to the function increment. Function increment
takes a pointer to int as argument.
- Variable at the address ptr is accessed and its value is incremented.
- Finally, incremented value is printed to screen. 


Point to Remember
- Call by reference is implemented indirectly by passing the address of a
variable to the function.


### Structure
Go language supports structures, which are a collection of multiple data types as a single entity. 

Example
```
package main

import "fmt"

type student struct {
	rollNo int
	name   string
}

func main() {
	stud := student{1, "Johnny"}
	fmt.Println(stud)
	fmt.Println("Student name ::", stud.name) // Accessing inner fields.
	pstud := &stud
	fmt.Println("Student name ::", pstud.name)   // Accessing inner fields.
	fmt.Println(student{rollNo: 2, name: "Ann"}) // Named initialization.
	fmt.Println(student{name: "Ann", rollNo: 2}) // Order does not matter.
	fmt.Println(student{name: "Alice"})          // Default initialization of rollNo.
}

```

Output
```
{1 Johnny}
Student name :: Johnny
Student name :: Johnny
{2 Ann}
{2 Ann}
{0 Alice}
```

### Methods
Go is an object-oriented language. However, it doesn't have any class keyword. We can associate fuctions directly by structures and other data types. 

Between func keyword and name of function, we add data type called 'receiver'. Once function is associated by a receiver, we can directly call function over structure using dot (.) operator. 

Example

```
package main

import "fmt"

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func main() {
	r := Rect{width: 10, height: 10}
	fmt.Println("Area: ", r.Area())
	fmt.Println("Perimeter: ", r.Perimeter())
	ptr := &Rect{width: 10, height: 5}
	fmt.Println("Area: ", ptr.Area())
	fmt.Println("Perimeter: ", ptr.Perimeter())
}
```

Output
```
Area:  100
Perimeter:  40
Area:  50
Perimeter:  30
```

Analysis:
- In the above program, we have Rect structure, which represent rectangle.
Rect has width and height fields.
- We associate two function Area() and Parameter() and associate them with
Rect data type.
- In the main() function we had created an instance of Rect, with width10
and height 10.
- Then Area() and Perimeter() functions are called using dot(.) operator
over Rect instance r.
- Then another instance of Rect is created and its address is stored in
pointer ptr.
- Again Area() and Perimeter() function is called using dot(.) operator but
this time we are calling the associated function over pointer. The go
language automatically convert the pointer to value and called the
associated functions.



There are 2 types of defining associated function of a data type.
1. Accessor function, which take receiver as value. Go passes a copy of the
instance this function (Just like pass by value.). Any change done over the
object is not reflected in the calling object.
The syntax of accessor function:
```
func (r <Receiver Data type>) <Function Name>(<Parameter List>) (<Return
List>)
```
2. Modifier function, which take receiver as pointer. Go passes actual
instance of the object to this function (Just like pass by pointer.) Any change
done over the object is reflected on the original object.
The syntax of modifier function:
```
func (r *<Receiver Data type>) <Function Name>(<Parameter List>) (<Return
List>)
```

Example
```
package main

import "fmt"

type MyInt int

func (data MyInt) increment1() {
	data = data + 1
}
func (data *MyInt) increment2() {
	*data = *data + 1
}
func main() {
	var data MyInt = 1
	fmt.Println("value before increment1() call :", data)
	data.increment1()
	fmt.Println("value after increment1() call :", data)
	data.increment2()
	fmt.Println("value after increment2() call :", data)
}
```


Output
```
value before increment1() call : 1
value after increment1() call : 1
value after increment2() call : 2
```

Analysis:
- Accessor function increment1() does changes on a local copy of the
object. Therefore, changes done are lost.
- Modifier function increment2() does changes on the actual instance so
changes done are preserved.






### Interface
Interfaces are defined as a set of methods. 

Syntax
```
Type <Interface name> interface {
<Method name> <Return type>
}
```


In Go, to implement an interface, an object just need to implement all methods of interface. When an object implement a particular interface, its object can be assigned to an interface type variable. 

Example

```
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}
func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}
func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 10}
	fmt.Println("Total Area: ", TotalArea(r, c))
	fmt.Println("Total Perimeter: ", TotalPerimeter(r, c))
}
```

Output
```
Total Area:  414.1592653589793
Total Perimeter:  102.83185307179586
```

Analysis

- A Shape interface is created which contain two methods Area() and
Perimeter().
- Rect and Circle implements Shape interface as they implement Area() and
Perimeter() methods.
- TotalArea() and TotalPerimeter() are two functions which expect list of
object of type Shape.


### Arrays
A collecition of variables of the same data type. 

Example
```
package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	count := len(arr)
	fmt.Println("Length of array", count)
}
```

Output
```
[0 0 0 0 0 0 0 0 0 0]
[0 1 2 3 4 5 6 7 8 9]
Length of array 10
```


### Slice

Go Array is a collection of variables of same data type. Arrays are fixed in legth and does not have any inbuilt method to increase their size. 

Go Slice is an abstraction over Array. It actually uses arrays as an underlying structure. The various operations are_
1. Inbuilt append() function is used to add the elements to a slice. If the size of
underlying array is not enough then automatically a new array is created and
content of the old array is copied to it.
2. The len() function returns the number of elements presents in the slice.
3. The cap() function returns the capacity of the underlying array of the slice.
4. The copy() function, the contents of a source slice are copied to a
destination slice.
5. Re-slices the slice, the syntax <Slice Name>[start : end] , It returns a slice
object containing the elements of base slice from index start to end1. Length of the new slice will be (end - start), and capacity will be cap
(base slice) - start.


To define a slice, you can declare it as an array without specifying its size.
Alternatively, you can use make function to create a slice.

Example
```
package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 17; i++ {
		s = append(s, i)
		PrintSlice(s)
	}
}
func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

```

Output
```
[1] :: len=1 cap=1 
[1 2] :: len=2 cap=2 
[1 2 3] :: len=3 cap=4 
[1 2 3 4] :: len=4 cap=4 
[1 2 3 4 5] :: len=5 cap=8 
[1 2 3 4 5 6] :: len=6 cap=8 
[1 2 3 4 5 6 7] :: len=7 cap=8 
[1 2 3 4 5 6 7 8] :: len=8 cap=8 
[1 2 3 4 5 6 7 8 9] :: len=9 cap=16 
[1 2 3 4 5 6 7 8 9 10] :: len=10 cap=16 
[1 2 3 4 5 6 7 8 9 10 11] :: len=11 cap=16 
[1 2 3 4 5 6 7 8 9 10 11 12] :: len=12 cap=16 
[1 2 3 4 5 6 7 8 9 10 11 12 13] :: len=13 cap=16 
[1 2 3 4 5 6 7 8 9 10 11 12 13 14] :: len=14 cap=16 
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15] :: len=15 cap=16 
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16] :: len=16 cap=16 
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17] :: len=17 cap=32 
```

Example
```
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintSlice(s) // [1 2 3 4 5 6 7 8 9 10] :: len=10 cap=10
	a := make([]int, 10)
	PrintSlice(a) // [0 0 0 0 0 0 0 0 0 0] :: len=10 cap=10
	b := make([]int, 0, 10)
	PrintSlice(b) // [] :: len=0 cap=10
	c := s[0:4]
	PrintSlice(c) // [1 2 3 4] :: len=4 cap=10
	d := c[2:5]
	PrintSlice(d) // [3 4 5] :: len=3 cap=8
}

func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

```


Output
```
[1 2 3 4 5 6 7 8 9 10] :: len=10 cap=10 
[0 0 0 0 0 0 0 0 0 0] :: len=10 cap=10 
[] :: len=0 cap=10 
[1 2 3 4] :: len=4 cap=10 
[3 4 5] :: len=3 cap=8 
```


### Map/Dictionary
A Map is a collection of Key-Value pairs. 
```
var <variable> map[<key datatype>]<value datatype>
```

Maps have to be initialized using make() before they can be used.
```
var <variable> map[<key datatype>]<value datatype> = make(map[<key datatype>]
<value datatype>)
```

or

```
<variable> := make(map[<key datatype>]<value datatype>)
```

Map Operations
1. Assignment: < variable>[<key>] = <value>
2. Delete: delete(<variable >, <key>)
3. Access: value, ok = < variable >[<key>] , the first value will have the
value of key in the map. If the key is not present in the map, it will return
zero value corresponding to the value data-type. The second argument
returns whether the map contains the key.


Example
```
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[ ", key, " -> ", val, " ]")
	}
	fmt.Println("Apple price:", m["Apple"])
	delete(m, "Apple")
	value, ok := m["Apple"]
	fmt.Println("Apple price:", value, "Present:", ok)
	value2, ok2 := m["Banana"]
	fmt.Println("Banana price:", value2, "Present:", ok2)
}

```

Outout
```
[ Apple -> 40 ][ Banana -> 30 ][ Mango -> 50 ]Apple price: 40
Apple price: 0 Present: false
Banana price: 30 Present: true
```



## Problem Solving - Array

- Write a method that will return the sum of all the elements of the integer list,
given list as an input argument.

```
package main

import "fmt"

func SumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

func main() {
	result := SumArray([]int{1, 2, 3, 4})
	fmt.Println(result)
}

```

Output
```
10
```


- Write a method, which will search a list for some given value.


First Approach (Sequential Search)
```
package main

import "fmt"

func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func main() {
	result := SequentialSearch([]int{1, 2, 3, 4}, 4)
	fmt.Println(result)
}
```

Output
```
true
```


Optimized Approch (Binary Search)
```
package main

import "fmt"

func BinarySearch(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			return true
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	result := BinarySearch([]int{1, 2, 3, 4}, 5)
	fmt.Println(result)
}
```
Output
```
false
```


- Given a list, you need to rotate its elements K number of times. For example, a
list [10,20,30,40,50,60] rotate by 2 positions to [30,40,50,60,10,20]

```
package main

import "fmt"

func RotateArray(data []int, k int) {
	ln := len(data)
	k = k % ln
	ReverseArray(data, 0, ln-1) //reverse full array
	ReverseArray(data, k, ln-1) //reverse the last part
	ReverseArray(data, 0, k-1)  //reverse the first part
	fmt.Println(data)

}
func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func main() {
	RotateArray([]int{-1, 2}, 3)
}


```

Output
```
[4 5 1 2 3]
```

- Given a list of positive and negative integers, find a contiguous subarray
whose sum (sum of elements) is maximum and return its sum.

```
func maxSubArray(nums []int) int {
    sumTotal := -10000
    nextPossibleSum := 0

    for i:=0; i<len(nums); i++{
        nextPossibleSum = nextPossibleSum+nums[i]
        if nextPossibleSum>sumTotal{
            sumTotal = nextPossibleSum
        }
        if nextPossibleSum<0{
            nextPossibleSum = 0
        }

    }

    return sumTotal
}
```


Here is [Part-1](2024-04-08-1-learn-go-dsa.md)

