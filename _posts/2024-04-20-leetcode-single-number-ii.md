---
title: "Leetcode - 137. Single Number II"
date: 2024-04-20
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Bit Manipulation]
---


# [Leetcode - 137. Single Number II](https://leetcode.com/problems/single-number-ii/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-single-number-ii.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-single-number-ii/)


- Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

  You must implement a solution with a linear runtime complexity and use only constant extra space.


Solution

```go
func singleNumber(nums []int) int {
    ones, twos := 0, 0

    for _, num := range nums {
        ones = (ones ^ num) &^ twos
        twos = (twos ^ num) &^ ones
    }

    return ones
}
```
