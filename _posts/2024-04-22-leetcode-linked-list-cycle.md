---
title: "Leetcode - 141. Linked List Cycle"
date: 2024-04-22
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table, Linked List, Two Pointers]
---


# [Leetcode - 141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-linked-list-cycle.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-linked-list-cycle/)


- Given head, the head of a linked list, determine if the linked list has a cycle in it.

  There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

  Return true if there is a cycle in the linked list. Otherwise, return false.


Solution

```
func hasCycle(head *ListNode) bool {

    if head == nil {
        return false
    }
    slowPointer := head
    fastPointer := head
    for fastPointer != nil && fastPointer.Next != nil {
        fastPointer = fastPointer.Next.Next
        slowPointer = slowPointer.Next

        if fastPointer == slowPointer {

            return true
        }
    }
    return false
}
```
