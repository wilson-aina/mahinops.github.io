---
title: "Leetcode - 1512. Number of Good Pairs"
date: 2024-04-15
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Hash Table, Math, Counting]
---


# [Leetcode - 1512. Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-number-of-good-pairs.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-number-of-good-pairs/)


- Given an array of integers nums, return the number of good pairs.

  A pair (i, j) is called good if nums[i] == nums[j] and i < j.


Solution 1

```go
func numIdenticalPairs(nums []int) int {
    pairCount := 0
	first, second := 0, 1

	for first=0; first<len(nums)-1; first++{
		for second=1; second<len(nums); second++{
			if nums[first]==nums[second] && first<second{
				pairCount++
			}
		}
	}
	return pairCount
}
```

Solution 2 (Optimized)

```go
func numIdenticalPairs(nums []int) int {
    pairCount := 0
	mapCount := make(map[int]int)

    for _, num := range nums{
        mapCount[num]++

        // every new occurrence of a number can be paired with every previous occurrence
		    // so if a number has already appeared 'p' times, we will have 'p-1' new pairs
        pairCount = pairCount+(mapCount[num]-1)
    }
	return pairCount
}
```

