---
title: "Leetcode - 100. Same Tree"
date: 2024-05-02
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Tree, Depth-First Search, DFS, Breadth-First Search, BFS, Binary Tree]
---


# [Leetcode - 100. Same Tree](https://leetcode.com/problems/same-tree/description/){:target="_blank"}
[![Hits](https://hits.sh/mahinops.github.io/posts/leetcode-same-tree.svg)](https://hits.sh/mahinops.github.io/posts/leetcode-same-tree/)


- Given the roots of two binary trees p and q, write a function to check if they are the same or not.

  Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Solution

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil{
        return true
    }

    if p==nil || q==nil{
        return false
    }
    
    if p.Val != q.Val{
        return false
    }

    left := isSameTree(p.Left, q.Left)
    right := isSameTree(p.Right, q.Right)

    return left && right
}
```
