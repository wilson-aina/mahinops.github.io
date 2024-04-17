---
title: "Leetcode - 287. Find the Duplicate Number"
date: 2024-04-18
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Two Pointers, Binary Search, Bit Manipulation]
---


# [Leetcode - 287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-find-the-duplicate-number.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-find-the-duplicate-number/)


- Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

  There is only one repeated number in nums, return this repeated number.

  You must solve the problem without modifying the array nums and uses only constant extra space.


Solution (Without Using Constant Space)

```go
func findDuplicate(nums []int) int {
    numMap := make(map[int]int)
    for _, num := range nums{
        _, ok := numMap[num]
        if !ok{
            numMap[num]++
        }else{
            return num
        }
    }
    return -1  
}
```

Solution (Constant Space/Best Practice)

```go
func findDuplicate(nums []int) int {
    i := 0

    for i<len(nums){
        indexWhereItShouldBe := nums[i]-1
        if nums[i]!=nums[indexWhereItShouldBe]{
            nums[i], nums[indexWhereItShouldBe] = nums[indexWhereItShouldBe], nums[i]
        }else{
            fmt.Println(i,indexWhereItShouldBe )
            if i != indexWhereItShouldBe {
                return nums[i]
            }
            i++
        }
    }

    fmt.Println(nums)
    return -1
}
```