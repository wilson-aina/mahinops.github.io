---
title: "Leetcode - 136. Single Number"
date: 2024-04-20
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Bit Manipulation]
---


# [Leetcode - 136. Single Number](https://leetcode.com/problems/single-number/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-single-number.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-single-number/)


- Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

  You must implement a solution with a linear runtime complexity and use only constant extra space.


Solution

```go
func singleNumber(nums []int) int {
    nonRepeatingNum := 0
    for _, num := range nums{
        // perform XOR operation that will emmit the number occuring double
        nonRepeatingNum^=num
    }
    return nonRepeatingNum 
}
```
