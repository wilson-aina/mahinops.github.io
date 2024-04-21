---
title: "Leetcode - 876. Middle of the Linked List"
date: 2024-04-22
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Linked List, Two Pointers]
---


# [Leetcode - 876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-middle-of-the-linked-list.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-middle-of-the-linked-list/)


- Given the head of a singly linked list, return the middle node of the linked list.

  If there are two middle nodes, return the second middle node.

Solution

```go
func middleNode(head *ListNode) *ListNode {
    if head==nil{
        return head
    }

    fastPointer := head
    slowPointer := head

    for fastPointer!=nil && fastPointer.Next!=nil{
        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next.Next
    }

    return slowPointer
}
```
