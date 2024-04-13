---
title: "Leetcode - 704. Binary Search"
date: 2024-04-14
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Binary Search]
---

# [Leetcode - 704. Binary Search](https://leetcode.com/problems/binary-search/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-binary-search.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-binary-search/)

- Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

  You must write an algorithm with O(log n) runtime complexity.

  
Solution
```
func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low<=high{
        mid := (low+high)/2
        fmt.Println(mid)
        if nums[mid]==target{
            return mid
        }else if target>nums[mid]{
            low = mid+1
        }else{
            high = mid-1
        }
    }
    return -1
}

```
