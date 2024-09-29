---
title: "Leetcode - 268. Missing Number"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table, Math, Binary Search, Bit Manipulation, Sorting]
---

# [Leetcode - 268. Missing Number](https://leetcode.com/problems/missing-number/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-missing-number.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-missing-number/)

- Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.



Solution
```go
func missingNumber(nums []int) int {
    ln := len(nums)
    sumTotal := (ln*(ln+1))/2
    numSum := 0

    for _, num := range nums{
        numSum+=num
    }
    return sumTotal-numSum
}

```
