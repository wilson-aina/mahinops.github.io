---
title: "Leetcode - 20. Valid Parentheses"
date: 2024-04-13
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, String, Stack]
---

# [Leetcode - 20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-valid-parentheses.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-valid-parentheses/)

- Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

    An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.



Solution
```go
func isValid(s string) bool {
    closeToOpenMap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    stack := []rune{}
    for _, char := range s {
        switch char{
            case '(', '[', '{':
                stack = append(stack, char)
            case ')', ']', '}':
                if len(stack)==0 || stack[len(stack)-1] != closeToOpenMap[char]{
                    return false
                }
                stack = stack[:len(stack)-1]
        }

    }
    return len(stack)==0
    
}
```
