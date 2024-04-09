---
title: "Leetcode - 217. Contains Duplicate"
date: 2024-04-10
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table, Sorting]
---

# [Leetcode - 217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-contains-duplicate.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-contains-duplicate/)

- Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Solution
```
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
