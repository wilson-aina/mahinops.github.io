---
title: "Leetcode - 15. 3Sum"
date: 2024-04-10
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Two Pointer, Sorting]
---

# [Leetcode - 15. 3Sum](https://leetcode.com/problems/3sum/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-3sum.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-3sum/)

- Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

    Notice that the solution set must not contain duplicate triplets.

Solution
```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for ind, val := range nums {
		if ind > 0 && val == nums[ind-1] {
			continue
		}
		left, right := ind+1, len(nums)-1
		for left < right {
			total := val + nums[left] + nums[right]
			if total > 0 {
				right -= 1
			} else if total < 0 {
				left += 1
			} else {
				res = append(res, []int{val, nums[left], nums[right]})
				left += 1
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}
	}
	return res
}


```
