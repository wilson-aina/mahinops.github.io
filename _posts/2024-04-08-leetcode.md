---
title: "Leetcode - 53. Maximum Subarray"
date: 2024-04-08
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Array, Divide and Conquer, Dynamic Programming, Programming, Algorithm]
---

# [Leetcode - 53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode/)

- Given an integer array nums, find the subarray with the largest sum, and return its sum.

Solution
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
