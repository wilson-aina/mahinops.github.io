---
title: "Leetcode - 206. Reverse Linked List"
date: 2024-04-17
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Linked List, Recursion]
---


# [Leetcode - 206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-reverse-linked-list.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-reverse-linked-list/)


- Given the head of a singly linked list, reverse the list, and return the reversed list.


Solution

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    current := head
    var prev, nextFlag *ListNode

    for current!=nil{
        nextFlag = current.Next
        current.Next = prev
        prev = current
        current = nextFlag

    }
    
    return prev
}
```
