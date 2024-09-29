---
title: "Leetcode - 11. Container With Most Water"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Array, Two Pointer, Greedy]
---

# [Leetcode - 11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-container-with-most-water.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-container-with-most-water/)

- You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

  Find two lines that together with the x-axis form a container, such that the container contains the most water.

  Return the maximum amount of water a container can store.

  Notice that you may not slant the container.

Solution
```go
func maxArea(height []int) int {
    leftH, rightH := 0, len(height)-1
    maxWater := 0

    for leftH<rightH{
        if height[leftH]<height[rightH]{
            area := (rightH-leftH)*height[leftH]
            if area>maxWater{
                maxWater = area
            }            
            leftH++
        }else{
            area := (rightH-leftH)*height[rightH]
            if area>maxWater{
                maxWater = area
            }
            rightH--
        }
    }
    return maxWater
}

```
