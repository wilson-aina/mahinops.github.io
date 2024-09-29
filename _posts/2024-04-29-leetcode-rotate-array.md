---
title: "Leetcode - 189. Rotate Array"
date: 2024-04-29
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Math, Two Pointers]
---


# [Leetcode - 189. Rotate Array](https://leetcode.com/problems/rotate-array/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-rotate-array.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-rotate-array/)


- Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.


Solution

```go
func rotate(nums []int, k int)  {
    ln := len(nums)
    if ln ==0{
        return
    }
    k = k%ln

    ReverseArray(nums, 0, ln-1)
    ReverseArray(nums, k, ln-1)
    ReverseArray(nums, 0, k-1)
}

func ReverseArray(data []int, start int, end int){
    for start<end{
        data[start], data[end] = data[end], data[start]
        start++
        end--
    }
}
```
