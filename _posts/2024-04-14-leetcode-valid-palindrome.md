---
title: "Leetcode - 125. Valid Palindrome"
date: 2024-04-14
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Two Pointer, String]
---


# [Leetcode - 125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-valid-palindrome.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-valid-palindrome/)


- A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

  Given a string s, return true if it is a palindrome, or false otherwise.


Solution
```go
func isPalindrome(s string) bool {
    firstInd, lastInd := 0, len(s)-1

    for firstInd<lastInd{
        for firstInd<lastInd && !isLetterOrDigit(rune(s[firstInd])){
            firstInd++
        }
        for firstInd<lastInd && !isLetterOrDigit(rune(s[lastInd])){
            lastInd--
        }

        if strings.ToLower(string(s[firstInd])) != strings.ToLower(string(s[lastInd])){
            return false
        }
        firstInd++
        lastInd--
    }
    return true
    
}


func isLetterOrDigit(r rune)bool{
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}
```
