---
title: "Leetcode - 41. First Missing Positive"
date: 2024-04-21
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table]
---


# [Leetcode - 41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-first-missing-positive.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-first-missing-positive/)


- Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

  You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.


Solution

`n will always belong to [1...len(nums)+1] inclusive`


```go
func firstMissingPositive(nums []int) int {
  i:=0

    // Rearrange the elements to place each positive integer at its correct index.
    // Negative numbers and numbers greater than the array size are ignored.
	for i < len(nums) {
    if nums[i]>0 && nums[i]<len(nums) && nums[i]!=nums[nums[i]-1]{
      nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
    }else{
      i++
    }
	}
  fmt.Println(nums)

  for i:=0; i<len(nums); i++{
    if nums[i] != i+1{
      return i+1
    }
  }

  return len(nums)+1
}
```
