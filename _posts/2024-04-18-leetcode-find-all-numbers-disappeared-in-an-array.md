---
title: "Leetcode - 448. Find All Numbers Disappeared in an Array"
date: 2024-04-18
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table]
---


# [Leetcode - 448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-find-all-numbers-disappeared-in-an-array.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-find-all-numbers-disappeared-in-an-array/)


- Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.



Solution

```go
func findDisappearedNumbers(nums []int) []int {
    i := 0

    for i<len(nums){
        indexOfCurrentNumber := nums[i]-1
        if nums[i]!=nums[indexOfCurrentNumber]{
            nums[i], nums[indexOfCurrentNumber] = nums[indexOfCurrentNumber], nums[i]
        }else{
            i++
        }
    }
    
    result := []int{}

    for ind, val := range nums{
        if (ind+1)!=val{
            result = append(result, ind+1)
        }
    }

    return result
}
```
