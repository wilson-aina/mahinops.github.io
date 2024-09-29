---
title: "Leetcode - 26. Remove Duplicates from Sorted Array"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Two Pointer]
---

# [Leetcode - 26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-remove-duplicates-from-sorted-array.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-remove-duplicates-from-sorted-array/)

- Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

  Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

  Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
  Return k.

Solution
```go
func removeDuplicates(nums []int) int {
    if len(nums)==0{
        return 0
    }
    nextUniqueIndex := 1

    for i:=1; i<len(nums); i++{
        if nums[i] !=nums[i-1]{
            nums[nextUniqueIndex] = nums[i]
            nextUniqueIndex++
        }
    }
    return nextUniqueIndex
}

```
