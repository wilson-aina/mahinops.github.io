---
title: "Leetcode - 66. Plus One"
date: 2024-04-12
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Math]
---

# [Leetcode - 66. Plus One](https://leetcode.com/problems/plus-one/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-plus-one.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-plus-one/)

- You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

	Increment the large integer by one and return the resulting array of digits.

Solution
```go
func plusOne(digits []int) []int {
    ln := len(digits)
    rem := 1
    for index:=ln-1; index>=0; index--{
        totalSumofLastIndex := digits[index]+1
        if totalSumofLastIndex<10{
            digits[index]=totalSumofLastIndex
            return digits
        }else{
            digits[index]=totalSumofLastIndex%10
            rem = totalSumofLastIndex/10

        }
    }
    return append([]int{rem}, digits...)
}

```
