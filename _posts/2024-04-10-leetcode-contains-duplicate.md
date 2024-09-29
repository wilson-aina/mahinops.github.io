---
title: "Leetcode - 217. Contains Duplicate"
date: 2024-04-10
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table, Sorting]
---

# [Leetcode - 217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-contains-duplicate.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-contains-duplicate/)

- Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Solution
```go
func containsDuplicate(nums []int) bool {
    numMap:= make(map[int]int)

    for _, j:=range(nums){
        _, ok:= numMap[j]
        if ok{
            return true
        }else{
            numMap[j]+=1
        }
    }
    return false
}
```


Another Approach 

```go
func containsDuplicate(nums []int) bool {
    sort.Ints(nums) // sort the slice
    // use a loop to compare each element with its next element
    for i := 0; i < len(nums)-1; i++ {
        // if any two elements are the same, return true
        if nums[i] == nums[i+1] {
            return true
        }
    }
    return false // if no duplicates are found, return false
}
```
