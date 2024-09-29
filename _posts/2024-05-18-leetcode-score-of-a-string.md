---
title: "Leetcode - 3110. Score of a String"
date: 2024-05-18
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, String]
---


# [Leetcode - 3110. Score of a String](https://leetcode.com/problems/score-of-a-string/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-score-of-a-string.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-score-of-a-string/)


- You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.

  Return the score of s.


Solution

```go
func scoreOfString(s string) int {
    sumTotal := 0

    for i:=1; i<len(s); i++{
        diff := int(math.Abs(float64(s[i]) - float64(s[i-1])))
        sumTotal += diff
    }

    return sumTotal
}
```
