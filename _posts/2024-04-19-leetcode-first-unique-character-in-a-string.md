---
title: "Leetcode - 387. First Unique Character in a String"
date: 2024-04-19
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table, String, Queue, Counting]
---


# [Leetcode - 387. First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-first-unique-character-in-a-string.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-first-unique-character-in-a-string/)


- Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.


Solution

```go
func firstUniqChar(s string) int {
    freqMap := make(map[rune]int)

    for _, char := range s{
        freqMap[char]++
    }

    for ind, char := range s{
        if freqMap[char]==1{
            return ind
        }
    }
    return -1
}
```
