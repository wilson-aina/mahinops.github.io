---
title: "Bit Manipulation"
date: 2024-04-20
categories: [Bit Manipulation]
tags: [Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Bit Manipulations]
---


# Bit Manipulation

Bit manipulation is a powerful technique often used in problem-solving, especially in competitive programming and algorithmic interviews. Here are some common topics related to bit manipulation:

Basic Operations:
1. Bitwise AND (&)
2. Bitwise OR (\|)
3. Bitwise XOR (^)
4. Bitwise NOT (~)
5. Left shift (<<)
6. Right shift (>>)


## Bitwise AND (&):
The bitwise AND operation returns 1 if both bits are 1, otherwise 0.

### Operation_

```go
package main

import "fmt"

func main() {
    result := 5 & 3
    fmt.Println(result)
}
```

### Explanations_
- We have two binary numbers: 5 (101 in binary) and 3 (011 in binary).
- Performing bitwise AND operation (&) on corresponding bits:

```bash
  101 
& 011 
--------
  001 -> which is 1 in decimal
```

Problems:
- Find if kth bit set or not set for a number.
  num = 7, k=2

Solutions_

1. Left shift k times of 1. 1<<k

```bash
num = 7 -> 111
k   = 2

Left shift 2 times of 1_
1 -> ...000000001
1<<2
Result -> ...0000000100

```

2. Do & operations with num. If the result if 0 means, kth bit is zero else 1.

```
Result -> ...00000100 (1<<2)
  Num=7 & ...00000111
  -----------------------
          ...00000100 -> kth bit result is 1. Means kth bit is 1
``` 

## Bitwise OR (\|)
The bitwise OR operation returns 1 if at least one of the bits is 1, otherwise 0.

### Operation_

```go
package main

import "fmt"

func main() {
    // Example of bitwise OR
    result := 5 | 3
    fmt.Println(result) // Output: 7
}
```

### Explanations_
```bash
Here, 5-> 101 and 3-> 011
  101 
| 011
--------
  111 -> Which is 7 in Decimal

```


## Bitwise XOR (^)
The bitwise XOR operation returns 1 if the bits are different, otherwise 0.

### Operation_

```go
package main

import "fmt"

func main() {
    // Example of bitwise XOR
    result := 5 ^ 3 
    fmt.Println(result) // Output: 6
}

```

### Explanations_
```bash
Here, 5-> 101 and 3-> 011
  101 
^ 011
--------
  110 -> Which is 6 in decimal.

```

## Bitwise NOT (~)
The bitwise NOT operation inverts each bit.

### Operation_

```go
package main

import "fmt"

func main() {
    // Example of bitwise NOT
    result := ^5 // ^101 = 010 (in two's complement representation)
    fmt.Println(result) // Output: -6
}

```

### Explanations_
```bash
Here, 5-> 101
^101 = 010 (in two's complement representation)

```

## Left shift (<<)
The left shift operation shifts the bits to the left by a specified number of positions.

### Operation_

```go
package main

import "fmt"

func main() {
    // Example of left shift
    result := 5 << 2 
    fmt.Println(result) // Output: 20
}

```

### Explanations_
```bash

5 -> 101 
Now,  101 << 2 = 10100

Output is 10 in decimal.

```



## Right shift (>>)
The right shift operation shifts the bits to the right by a specified number of positions.

### Operation_

```go
package main

import "fmt"

func main() {
    // Example of right shift
    result := 10 >> 1 
    fmt.Println(result) // Output: 5
}


```

### Explanations_
```bash
10 -> 1010 
Now,  1010 >> 1 = 0101

Output is 5 in decimal.
```


