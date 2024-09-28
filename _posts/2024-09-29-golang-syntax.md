---
title: "Golang Syntax"
date: 2024-09-29
categories: [Golang]
tags: [Golang, Learning Golang]
---

## Golang
Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. 

## Basic Data Types and Operations

### Numeric Types in Go

1. Integers:
- **Signed Integers**(can represent both positive and negative values.): `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned Integers: `uint`, `uint8` (alias byte), `uint16`, `uint32`, `uint64`

  The `int` and `uint` types are platform-dependent and will be either `32-bit` or `64-bit` based on your system architecture.

2. Floating-point Numbers:
- `float32` and `float64` (commonly used for decimal numbers)

3. Other Numeric Types:
- `uintptr`: Used for storing pointer data

### Basic Arithmetic Operations

```go
package main

import "fmt"

func main() {
    // Integer operations
    a := 10
    b := 3

    fmt.Println("Addition:", a+b)        // Output: 13
    fmt.Println("Subtraction:", a-b)     // Output: 7
    fmt.Println("Multiplication:", a*b)  // Output: 30
    fmt.Println("Division:", a/b)        // Output: 3 (integer division)
    fmt.Println("Modulus:", a%b)         // Output: 1

    // Floating-point operations
    x := 10.5
    y := 3.2

    fmt.Println("Float Addition:", x+y)       // Output: 13.7
    fmt.Println("Float Subtraction:", x-y)    // Output: 7.3
    fmt.Println("Float Multiplication:", x*y) // Output: 33.6
    fmt.Println("Float Division:", x/y)       // Output: 3.28125

    // Complex numbers
    c1 := complex(2, 3) // 2 + 3i
    c2 := complex(1, 4) // 1 + 4i

    fmt.Println("Complex Addition:", c1+c2)       // Output: (3+7i)
    fmt.Println("Complex Subtraction:", c1-c2)    // Output: (1-1i)
    fmt.Println("Complex Multiplication:", c1*c2) // Output: (-10+11i)
    fmt.Println("Complex Division:", c1/c2)       // Output: (0.8235294117647058-0.29411764705882354i)
}
```


### Numeric Operations in Go

Advanced Math Operations Using the math Package
- Square Root: `math.Sqrt(x)`
- Exponentiation: `math.Pow(x, y)`
- Trigonometric functions: `math.Sin(x)`, `math.Cos(x)`, `math.Tan(x)`
- Absolute Value: `math.Abs(x)`




### Map in Go

1. Basic Syntax `map[KeyType]ValueType`
2. Declaring an empty map using make: `numbers := make(map[string]int)`
3. Declaring and initializing a map with values:

    ```go
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    ```
4. Adding or updating elements:

    ```go
    numbers["one"] = 1
    numbers["two"] = 2
    ```
5. Retrieving elements:

    ```go
    value := numbers["one"]
    fmt.Println(value) // Output: 1
    ```
6. delete(numbers, "one"): `delete(numbers, "one")`

    **Note:**

    ```go
    numbers := make(map[string]int) //This approach uses the make function to create an empty map.
    numbers := map[string]int{} //This approach directly initializes an empty map using a map literal.
    ```



### String in Go
1. Concatenation

    ```go
    str1 := "Hello"
    str2 := "World"
    result := str1 + " " + str2
    fmt.Println(result) // Output: Hello World
    ```
2. String Length

    ```go
    str := "Golang"
    length := len(str)
    fmt.Println(length) // Output: 6
    ```

3. Accessing Characters

    ```go
    str := "Golang"
    char := str[0]
    fmt.Println(char)        // Output: 71 (ASCII code for 'G')
    fmt.Println(string(char)) // Output: G
    ```

4. Substring

    ```go
    str := "Golang"
    substring := str[1:4]
    fmt.Println(substring)
    ```

5. Checking Substrings

    ```go
    import "strings"

    str := "Hello, Golang!"
    contains := strings.Contains(str, "Golang")
    fmt.Println(contains) // Output: true
    ```

6. Replacing Substrings

    ```go
    import "strings"

    str := "Hello, World!"
    newStr := strings.Replace(str, "World", "Golang", 1)
    fmt.Println(newStr) // Output: Hello, Golang!
    ```

7. Splitting a String

    ```go
    import "strings"

    str := "apple,banana,cherry"
    fruits := strings.Split(str, ",")
    fmt.Println(fruits) // Output: [apple banana cherry]

    ```

8. Trimming Whitespaces

    ```go
    import "strings"

    str := "  Hello, Golang!  "
    trimmed := strings.TrimSpace(str)
    fmt.Println(trimmed) // Output: Hello, Golang!
    ```

9. Converting Case

    ```go
    import "strings"

    str := "GoLang"
    lower := strings.ToLower(str)
    upper := strings.ToUpper(str)
    fmt.Println(lower) // Output: golang
    fmt.Println(upper) // Output: GOLANG
    ```

10. Comparing Strings

    ```go
    import "strings"

    str1 := "hello"
    str2 := "world"
    equal := str1 == str2
    fmt.Println(equal) // Output: false

    comparison := strings.Compare(str1, str2)
    fmt.Println(comparison) // Output: -1 (str1 is less than str2)
    ```


## Array in Go

1. With a specified length:

    ```go
    var numbers [5]int // An array of 5 integers
    ```

2. Using the shorthand syntax:

    ```go
    numbers := [5]int{1, 2, 3, 4, 5} // An array with initial values
    ```

3. Using ... to let the compiler determine the length:

    ```go
    numbers := [...]int{1, 2, 3, 4, 5} // The length is automatically determined
    ```

4. Accessing and Modifying Array Elements

    ```go
    package main

    import "fmt"

    func main() {
        // Declare and initialize an array
        numbers := [5]int{1, 2, 3, 4, 5}

        // Accessing elements
        fmt.Println("First element:", numbers[0]) // Output: 1
        fmt.Println("Third element:", numbers[2])  // Output: 3

        // Modifying elements
        numbers[0] = 10
        fmt.Println("Modified first element:", numbers[0]) // Output: 10

        // Print the entire array
        fmt.Println("Array:", numbers) // Output: [10 2 3 4 5]
    }
    ```


4. Iterating Over an Array

    **for loop**

    ```go
    for i := 0; i < len(numbers); i++ {
        fmt.Println("Element at index", i, ":", numbers[i])
    }
    ```

    **for range**

    ```go
    for index, value := range numbers {
        fmt.Println("Element at index", index, ":", value)
    }
    ```



## Slices in Go

1. Declaring and Initializing Slices

    **Using `make`:**

    ```go
    ```

    **Using a slice literal:**

    ```go
    slice := []int{1, 2, 3, 4, 5} // Initializes a slice with values
    ```

    **Creating a slice from an array:**

    ```go
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:4] // Creates a slice from index 1 to 3
    ```

2. Accessing and Modifying Slice Elements

    ```go
    package main

    import "fmt"

    func main() {
        // Initialize a slice
        slice := []int{1, 2, 3, 4, 5}

        // Accessing elements
        fmt.Println("First element:", slice[0]) // Output: 1
        fmt.Println("Third element:", slice[2])  // Output: 3

        // Modifying elements
        slice[0] = 10
        fmt.Println("Modified first element:", slice[0]) // Output: 10

        // Print the entire slice
        fmt.Println("Slice:", slice) // Output: [10 2 3 4 5]
    }

    ```

3. Appending to a Slice

    ```go
    package main

    import "fmt"

    func main() {
        slice := []int{1, 2, 3}

        // Appending elements
        slice = append(slice, 4) // Appending a single element
        slice = append(slice, 5, 6) // Appending multiple elements

        fmt.Println("After appending:", slice) // Output: [1 2 3 4 5 6]
    }

    ```

4. Slicing a Slice

    ```go
    package main

    import "fmt"

    func main() {
        slice := []int{1, 2, 3, 4, 5}

        // Creating a new slice from an existing slice
        newSlice := slice[1:4] // Contains elements at index 1, 2, and 3

        fmt.Println("New slice:", newSlice) // Output: [2 3 4]
    }

    ```

5. Length and Capacity

    ```go
    package main

    import "fmt"

    func main() {
        slice := []int{1, 2, 3}

        fmt.Println("Length:", len(slice)) // Output: 3
        fmt.Println("Capacity:", cap(slice)) // Output: 3 (initial capacity)
        
        slice = append(slice, 4, 5)
        
        fmt.Println("New Length:", len(slice)) // Output: 5
        fmt.Println("New Capacity:", cap(slice)) // Output may increase (e.g., 6, 8, etc.)
    }

    ```

6. Iterating Over an Slice

    **Using a traditional `for` loop:**

    ```go
    for i := 0; i < len(slice); i++ {
        fmt.Println("Element at index", i, ":", slice[i])
    }

    ```

    **Using `for range`:**

    ```go
    for index, value := range slice {
        fmt.Println("Element at index", index, ":", value)
    }

    ```









































































