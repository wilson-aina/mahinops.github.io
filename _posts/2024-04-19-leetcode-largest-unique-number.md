---
title: "Leetcode - Largest Unique Number(Premium)"
date: 2024-04-19
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table]
---


# [Leetcode - Largest Unique Number(Premium)](https://leetcode.com/problems/largest-unique-number/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-largest-unique-number.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-largest-unique-number/)


- Given an array of integers, identify the highest value that appears only once in the array. If no such number exists, return -1.


Solution

```go
type Solution struct{}

func (sol Solution) largestUniqueNumber(A []int) int {
	freqMap := make(map[int]int)
	largestVal := -1
	for _, val := range A{
		freqMap[val]++
	}

	for _, val := range A{
		if freqMap[val]==1 && val>largestVal{
			largestVal= val
		}
	}
	return largestVal
}

```
