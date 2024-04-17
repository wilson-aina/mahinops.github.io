---
title: "In-Place Cyclic Sort"
date: 2024-04-18
categories: [Sorting]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Sorting]
---


# In-Place Cyclic Sort

As we know, the input array contains numbers from the range 1 to n. We can use this fact to devise an efficient way to sort the numbers. Since all numbers are unique, we can try placing each number at its correct place, i.e., placing 1 at index ‘0’, placing 2 at index ‘1’, and so on.

To place a number (or an object in general) at its correct index, we first need to find that number. If we first find a number and then place it at its correct place, it will take us , which is not acceptable as mentioned in the problem statement.

Instead, what if we iterate the array one number at a time, and if the current number we are iterating is not at the correct index, we swap it with the number at its correct index. This way, we will go through all numbers and place them at their correct indices, hence, sorting the whole array.

<div style="text-align:center">
    <img src="static/cyclic-sort1.png" alt="Cyclic Sort-1">
    <img src="static/cyclic-sort2.png" alt="Cyclic Sort-2">
    <p>Fig: Cyclic Sort</p>
</div>

```go
package main

import "fmt"

// in place sorting
func cyclicSort(nums []int) []int {
	i := 0

	// technique: check number and check if they are in right index

	for i < len(nums) {
		idxOfCurrentElement := nums[i] - 1 // index of the current element where it should be placed

		if nums[i] != nums[idxOfCurrentElement] { // Check if the current element is not in its correct position.
			nums[i], nums[idxOfCurrentElement] = nums[idxOfCurrentElement], nums[i] // Swap the current element with the one at its correct position.
		} else {
			i++
		}
	}

	return nums
}

func main() {
	// unsorted from 1 to n
	arr := []int{1, 3, 4, 2, 5}

	result := cyclicSort(arr)
	fmt.Println("Sorted array is", result)

}

```

Output

```bash
Sorted array is [1 2 3 4 5]
```

