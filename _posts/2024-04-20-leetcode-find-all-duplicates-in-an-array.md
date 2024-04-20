---
title: "Leetcode - 442. Find All Duplicates in an Array"
date: 2024-04-20
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table]
---


# [Leetcode - 442. Find All Duplicates in an Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-find-all-duplicates-in-an-array.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-find-all-duplicates-in-an-array/)


- Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.

  You must write an algorithm that runs in O(n) time and uses only constant extra space.


Solution

```
func findDuplicates(nums []int) []int {
    i := 0

    for i<len(nums){
        indexWhereItShouldBe := nums[i]-1
        if nums[i]!=nums[indexWhereItShouldBe]{
            nums[i], nums[indexWhereItShouldBe] = nums[indexWhereItShouldBe], nums[i]
        }else{
            i++
        }
    }

    duplicateNumbers := make([]int, 0)
	for i = 0; i < len(nums); i++ {
		if nums[i] != i+1 { 
			duplicateNumbers = append(duplicateNumbers, nums[i]) 
		}
	}
	return duplicateNumbers
    
}
```
