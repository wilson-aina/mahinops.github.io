---
title: "Leetcode - 142. Linked List Cycle II"
date: 2024-04-23
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table, Linked List, Two Pointers]
---


# [Leetcode - 142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-linked-list-cycle-ii.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-linked-list-cycle-ii/)


- Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

  There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

  Do not modify the linked list.


Solution

```
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    flag := head
    slowPointer, fastPointer := head, head

    for fastPointer!=nil && fastPointer.Next!=nil{
        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next.Next

        if slowPointer==fastPointer{
            result:=CycleResult(flag, slowPointer)
            return result
        }
    }
    return nil
}

func CycleResult(head, slow *ListNode)(*ListNode){
    p := head
    for{
        if p==slow{
            return slow
        }
        p = p.Next
        slow = slow.Next
    }
}
```
