---
title: "Leetcode - 1. Two Sum"
date: 2024-04-08
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table]
---

# [Leetcode - 1. Two Sum](https://leetcode.com/problems/two-sum/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-two-sum.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-two-sum/)

- Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

    You may assume that each input would have exactly one solution, and you may not use the same element twice.

    You can return the answer in any order.

Solution
```go
func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num:= range nums{
        numberToFind := target-num
        if hashMapIndex, ok:= hashMap[numberToFind]; ok{
            return []int{hashMapIndex, i}
        }
        hashMap[num]=i
    }
    return []int{}
}

```
