---
title: "Leetcode - 242. Valid Anagram"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table, String, Sorting]
---

# [Leetcode - 242. Valid Anagram](https://leetcode.com/problems/valid-anagram/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-valid-anagram.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-valid-anagram/)

- Given two strings s and t, return true if t is an anagram of s, and false otherwise.

    An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Solution
```go
func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    charCount := make(map[rune]int) //rune used to handle Unicode character correctly
    
    for _, char := range s{
        charCount[char]++
    }
    
    for _, char := range t{
        charCount[char]--
        if charCount[char]<0{
            return false
        }
    }
    return true
}

```
