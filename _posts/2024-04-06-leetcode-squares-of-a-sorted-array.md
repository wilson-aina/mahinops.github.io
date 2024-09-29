---
title: "Leetcode - 977. Squares of a Sorted Array"
date: 2024-04-06
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Two Pointer, Sorting]
---

# [Leetcode - 977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-squares-of-a-sorted-array.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-squares-of-a-sorted-array/)

- Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.


Solution
```go
func sortedSquares(nums []int) []int {
    ln := len(nums)
    sortedSquared := make([]int, ln)
    highestIndex := ln-1
    
    left, right := 0, ln-1

    for left<=right{
        squaredLeft := nums[left]*nums[left]
        squaredRight := nums[right]*nums[right]

        if squaredLeft>squaredRight{
            sortedSquared[highestIndex] = squaredLeft
            highestIndex--
            left++
        }else{
            sortedSquared[highestIndex] = squaredRight
            highestIndex--
            right--
        }
    }
    return sortedSquared
}

```
