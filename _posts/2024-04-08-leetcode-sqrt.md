---
title: "Leetcode - 69. Sqrt(x)"
date: 2024-04-08
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Math, Binary Search]
---

# [Leetcode - 69. Sqrt(x)](https://leetcode.com/problems/sqrtx/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-sqrt.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-sqrt/)

- Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

    You must not use any built-in exponent function or operator.

    For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Solution
```
func mySqrt(x int) int {
    if x<2{
        return x
    }

    minVal, maxVal := 1, x/2

    for minVal<=maxVal{
        midVal := (minVal+maxVal)/2

        midSquareVal := midVal*midVal

        if midSquareVal == x{
            return midVal
        }
        if midSquareVal>x{
            maxVal = midVal-1
        }else{
            minVal = midVal+1
        }
    }
    return maxVal
}

```
