---
title: "Leetcode - 167. Two Sum II - Input Array Is Sorted"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Two Pointer, Sorting]
---

# [Leetcode - 167. Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-two-sum-ii-input-array-is-sorted.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-two-sum-ii-input-array-is-sorted/)

- Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

    Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

    The tests are generated such that there is exactly one solution. You may not use the same element twice.

    Your solution must use only constant extra space.


Solution
```
func twoSum(numbers []int, target int) []int {
    firstInd, lastInd := 0, len(numbers)-1

    for firstInd<lastInd{
        if numbers[lastInd]+numbers[firstInd]==target{
            return []int{firstInd+1, lastInd+1}
        }
        if numbers[lastInd]+numbers[firstInd]>target{
            lastInd--
        }else{
            firstInd++
        }
    }
    return []int{0}
}
```
